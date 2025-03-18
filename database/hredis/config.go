package hredis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Option func(*connection)

func WithAddr(addr string) Option {
	return func(c *connection) {
		c.addr = addr
	}
}

func WithDatabase(db int) Option {
	return func(c *connection) {
		c.db = db
	}
}

type connection struct {
	addr string
	db   int
	cli  *redis.Client
}

func NewConnection(ctx context.Context, opts ...Option) {
	onceRedis.Do(func() {
		redisConn = &connection{addr: defaultAddr}
		for _, opt := range opts {
			opt(redisConn)
		}

		redisConn.conn(ctx)
	})
}

func (c *connection) conn(ctx context.Context) *connection {
	c.cli = redis.NewClient(&redis.Options{
		Addr: c.addr,
		DB:   c.db,
	})

	_, err := c.cli.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return c
}
