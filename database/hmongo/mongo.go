package hmongo

import (
	"context"
	"sync"
)

var (
	onceMongo       sync.Once
	mongoConn       *connection
	mongoDefaultUrl = "mongodb://localhost:27017"
)

func GetConnection(ctx context.Context, collection string) *connection {
	mongoConn.collection = mongoConn.conn(ctx).withCollection(ctx, collection)
	return mongoConn
}
