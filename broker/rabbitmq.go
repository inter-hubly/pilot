package broker

import (
	"sync"

	"github.com/streadway/amqp"
)

var (
	rabbitOnce      sync.Once
	rabbitConn      *rabbitMQ
	rabbitMqDefault = "amqp://guest:guest@localhost:5672/"
)

type Connection interface {
	Publish(routingKey string, body []byte) error
	Consume(queue string, consumeFunc func(value amqp.Delivery))
	QueueBind(queue, routingKey, exchange string) error
	Close()
}

func GetConnection() *rabbitMQ {
	return rabbitConn
}
