package hmongo

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	onceMongo       sync.Once
	mongoConn       *connection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

func GetConnection(ctx context.Context) *connection {
	return mongoConn
}

type NoSqlConn interface {
	GetCollection(ctx context.Context, collection string) *mongo.Collection
}

func (c *connection) GetCollection(ctx context.Context, collection string) *mongo.Collection {
	return c.mongo.Collection(collection)
}
