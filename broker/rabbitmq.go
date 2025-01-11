package broker

import (
	"context"
	"sync"

	"github.com/streadway/amqp"
)

var (
	rabbitOnce      sync.Once
	rabbitConn      *rabbitMQ
	rabbitMqDefault = "amqp://guest:guest@localhost:5672/"
)

type Connection interface {
	Publish(ctx context.Context, routingKey string, body []byte) error
	Consume(ctx context.Context, queue string, consumeFunc func(value amqp.Delivery))
	QueueBind(ctx context.Context, queuesBind ...*queueBinding) error
	Close(ctx context.Context)
}

func GetConnection() *rabbitMQ {
	return rabbitConn
}
