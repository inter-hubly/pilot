package hmongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Option func(*connection)

func WithUrl(url string) Option {
	return func(a *connection) {
		a.url = url
	}
}

func WithDatabase(db string) Option {
	return func(a *connection) {
		a.database = db
	}
}

type connection struct {
	url      string
	database string
	mongo    *mongo.Database
}

func NewConnection(ctx context.Context, opts ...Option) {
	onceMongo.Do(func() {
		mongoConn = &connection{}
		mongoConn.url = mongoDefaultUrl
		for _, opt := range opts {
			opt(mongoConn)
		}

		mongoConn.conn(ctx)
	})
}

func (c *connection) conn(ctx context.Context) *connection {
	clientOptions := options.Client().ApplyURI(c.url)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	c.mongo = client.Database(c.database)
	return c
}

func (c *connection) withCollection(ctx context.Context, collection string) *mongo.Collection {
	return c.mongo.Collection(collection)
}
