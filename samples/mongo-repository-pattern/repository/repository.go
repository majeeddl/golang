package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepositoryStruct[T any] struct {
	collection *mongo.Collection
}

func MongoRepository[T any](client *mongo.Client, databaseName string, collectionName string) *MongoRepositoryStruct[T] {
	collection := client.Database(databaseName).Collection(collectionName)
	return &MongoRepositoryStruct[T]{collection}
}

func (r *MongoRepositoryStruct[T]) Find(ctx context.Context, filter bson.M) ([]T, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *MongoRepositoryStruct[T]) FindOne(ctx context.Context, filter bson.M) (T, error) {

	var result T
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (r *MongoRepositoryStruct[T]) InsertOne(ctx context.Context, document T) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(ctx, document)
	return result, err
}
