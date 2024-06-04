package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type IRepository[T any] interface {
	Find(ctx context.Context, filter bson.M) ([]T, error)
	FindOne(ctx context.Context, filter bson.M) (T, error)
	InsertOne(ctx context.Context, document T) error
}
