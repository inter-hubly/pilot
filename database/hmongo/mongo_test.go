package hmongo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	NewConnection(ctx, WithUrl("mongodb://localhost:27017"), WithDatabase("skadi"))

	test := struct {
		Id   string `bson:"_id"`
		Name string `bson:"name"`
	}{}
	GetConnection(ctx).GetCollection(ctx, "user").FindOne(ctx, bson.M{"_id": "658759f4abfaf2451945f168"}).Decode(&test)
}
