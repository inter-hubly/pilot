package hmongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/inter-hubly/pilot/domain/valueobject"
	"github.com/inter-hubly/pilot/hlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSqlConn interface {
	FindById(ctx context.Context, id string, object any) error
	FindByIdWithProjection(ctx context.Context, idStr string, fields ...string) (map[string]interface{}, error)
	FindByFieldWithProjection(ctx context.Context, idStr valueobject.Pair[string, string], fields ...string) (map[string]interface{}, error)
}

func (c *connection) FindById(ctx context.Context, idStr string, object any) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		hlog.Error(ctx, "NoSqlConn.FindById", fmt.Sprintf("Invalid ObjectID: %v", err))
		return fmt.Errorf("invalid object ID: %w", err)
	}

	err = c.collection.FindOne(ctx, bson.M{"_id": id}).Decode(object)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			hlog.Warn(ctx, "NoSqlConn.FindById", "Document not found")
			return fmt.Errorf("document not found: %w", err)
		}
		hlog.Error(ctx, "NoSqlConn.FindById", fmt.Sprintf("Error decoding object: %v", err))
		return fmt.Errorf("failed to decode object: %w", err)
	}

	return nil
}

func (c *connection) FindByIdWithProjection(ctx context.Context, idStr string, fields ...string) (map[string]interface{}, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		hlog.Error(ctx, "NoSqlConn.FindByIdWithProjection", fmt.Sprintf("Invalid ObjectID: %v", err))
	}

	projection := bson.M{}
	for _, field := range fields {
		projection[field] = 1
	}

	opts := options.FindOne().SetProjection(projection)

	result := make(map[string]interface{})
	err = c.collection.FindOne(ctx, bson.M{"_id": id}, opts).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			hlog.Warn(ctx, "NoSqlConn.FindById", "Document not found")
			return nil, fmt.Errorf("document not found: %w", err)
		}
		hlog.Error(ctx, "NoSqlConn.FindById", fmt.Sprintf("Error decoding object: %v", err))
		return nil, fmt.Errorf("failed to decode object: %w", err)
	}
	return result, nil
}

func (c *connection) FindByFieldWithProjection(ctx context.Context, pairValue valueobject.Pair[string, string], fields ...string) (map[string]interface{}, error) {
	projection := bson.M{}
	for _, field := range fields {
		projection[field] = 1
	}

	opts := options.FindOne().SetProjection(projection)

	result := make(map[string]interface{})

	if err := c.collection.FindOne(ctx, bson.M{pairValue.Key: pairValue.Value}, opts).Decode(result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			hlog.Warn(ctx, "NoSqlConn.FindById", "Document not found")
			return nil, fmt.Errorf("document not found: %w", err)
		}
		hlog.Error(ctx, "NoSqlConn.FindById", fmt.Sprintf("Error decoding object: %v", err))
		return nil, fmt.Errorf("failed to decode object: %w", err)
	}
	return result, nil
}
