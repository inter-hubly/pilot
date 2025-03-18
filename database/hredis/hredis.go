package hredis

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisConn   *connection
	onceRedis   sync.Once
	defaultAddr = "localhost:6379"
)

func GetConnection(ctx context.Context) *connection {
	return redisConn
}

type RedisConn interface {
	GetClient(ctx context.Context) *redis.Client
}

func (c *connection) GetClient(ctx context.Context) *redis.Client {
	return c.cli
}
