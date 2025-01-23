package hmongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	NewConnection(ctx, WithUrl("mongodb://localhost:27017"), WithDatabase("skadi"))

	test := struct {
		Id   string `bson:"_id"`
		Name string `bson:"name"`
	}{}
	GetConnection(ctx, "user").FindById(ctx, "658759f4abfaf2451945f168", &test)
}

func TestMongoWithProjection(t *testing.T) {
	ctx := context.Background()
	NewConnection(ctx, WithUrl("mongodb://localhost:27017"), WithDatabase("skadi"))

	projection, err := GetConnection(ctx, "user").FindByIdWithProjection(ctx, "658759f4abfaf2451945f168", "_id")
	assert.Nil(t, err)
	assert.Nil(t, projection["name"])
}
