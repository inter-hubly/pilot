package broker

import (
	"fmt"
	"log"

	"github.com/inter-hubly/pilot/hlog"
	"github.com/streadway/amqp"
)

// Option defines a functional option for configuring RabbitMQ.
type Option func(*rabbitMQ)

// WithURL allows customization of the RabbitMQ connection URL.
func WithURL(url string) Option {
	return func(r *rabbitMQ) {
		r.url = url
	}
}

type queueBinding struct {
	QueueName  string
	RoutingKey string
	Exchange   string
}

func NewQueueBinding(queueName, routingKey, exchangeName string) *queueBinding {
	return &queueBinding{
		QueueName:  queueName,
		RoutingKey: routingKey,
		Exchange:   exchangeName,
	}
}

// RabbitMQ represents the RabbitMQ connection and channel.
type rabbitMQ struct {
	url      string
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
}

// NewRabbitMQ initializes a singleton RabbitMQ connection.
func NewRabbitMQ(exchangeName, exchangeType string, opts ...Option) {
	rabbitOnce.Do(func() {
		rabbitConn = &rabbitMQ{
			url: rabbitMqDefault,
		}

		// Apply functional options
		for _, opt := range opts {
			opt(rabbitConn)
		}

		// Establish connection and channel
		if err := rabbitConn.connect(); err != nil {
			log.Fatalf("Failed to initialize RabbitMQ: %v", err)
		}

		if err := rabbitConn.declareExchange(exchangeName, exchangeType); err != nil {
			log.Fatalf("Failed to declare RabbitMQ exchange: %v", err)
		}
	})
}

// connect establishes the connection and opens a channel.
func (r *rabbitMQ) connect() error {
	var err error

	// Connect to RabbitMQ
	r.conn, err = amqp.Dial(r.url)
	if err != nil {
		hlog.Error("RabbitMQ.connect", fmt.Sprintf("failed to connect to RabbitMQ: %s", err))
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Create a channel
	r.channel, err = r.conn.Channel()
	if err != nil {
		hlog.Error("RabbitMQ.connect", fmt.Sprintf("failed to open a channel: %s", err))
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	hlog.Info("RabbitMQ.connect", "RabbitMQ connected successfully")
	return nil
}

// declareExchange declares an exchange.
func (r *rabbitMQ) declareExchange(exchangeName, exchangeType string) error {
	err := r.channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error("rabbitMQ.declareExchange", fmt.Sprintf("failed to declare exchange: %s", err))
	}

	r.exchange = exchangeName
	hlog.Info("rabbitMQ.declareExchange", fmt.Sprintf("exchange %s declared successfully", exchangeName))
	return nil
}

// Publish sends a message to the exchange.
func (r *rabbitMQ) Publish(routingKey string, body []byte) error {
	err := r.channel.Publish(
		r.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Message published to exchange %s with routing key %s", r.exchange, routingKey)
	return nil
}

func (r *rabbitMQ) Consume(queue string, consumeFunc func(amqp.Delivery)) {
	r.queueDeclare(queue)
	msgs, err := r.channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error("rabbitMQ.Consume", "failed to consume message: ", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			hlog.Info("rabbitMQ.Consume", fmt.Sprintf("Fila %s: Recebida uma mensagem: %s", queue, d.Body))
			consumeFunc(d)
		}
	}()
	<-forever
}

// Close closes the connection and channel.
func (r *rabbitMQ) Close() {
	if r.channel != nil {
		_ = r.channel.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
	log.Println("RabbitMQ connection closed")
}

func (r *rabbitMQ) QueueBind(queuesBind ...*queueBinding) error {
	for _, queue := range queuesBind {
		if err := r.channel.QueueBind(
			queue.QueueName,
			queue.RoutingKey,
			queue.Exchange,
			false,
			nil,
		); err != nil {
			return err
		}
	}
	return nil
}

func (r *rabbitMQ) queueDeclare(queueName string) {
	_, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		hlog.Error("rabbitMQ.queueDeclare", fmt.Sprintf("failed to declare a queue: %s", err))
	}
}
