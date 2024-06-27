package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepositoryConfig struct {
	Ctx      context.Context
	Database *mongo.Database
}

type MongoRepository[T any] struct {
	Ctx        context.Context
	Database   *mongo.Database
	Collection string
}

func (r *MongoRepository[T]) Create(entity T) (T, error) {
	return entity, nil
}

func (r *MongoRepository[T]) FindAll() ([]T, error) {
	var result []T
	cur, err := r.Database.Collection(r.Collection).Find(r.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(r.Ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository[T]) FindByID(id string) (T, error) {
	var newEntity T
	return newEntity, nil
}

func (r *MongoRepository[T]) Update(id string, entity T) (T, error) {
	return entity, nil
}

func (r *MongoRepository[T]) Delete(id string) error {
	return nil
}
