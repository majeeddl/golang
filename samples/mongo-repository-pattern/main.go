package main

import (
	"context"
	"fmt"
	"mongorepositorypattern/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Mocking the mongo client and collection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	type User struct {
		ID   primitive.ObjectID `bson:"_id"`
		Name string             `bson:"name"`
	}
	// Creating the repository
	userRepository := repository.MongoRepository[User](client, "test", "users")

	// Mocking the collection's InsertOne method
	// collection.On("InsertOne", ctx, bson.M{"key": "value"}).Return(nil, nil)

	// Testing the InsertOne method
	_, err = userRepository.InsertOne(ctx, User{ID: primitive.NewObjectID(), Name: "masoumeh"})

	if err != nil {
		panic(err)
	}

	// assert.Nil(t, err)

	// Mocking the collection's FindOne method
	// userRepository.On("FindOne", ctx, bson.M{"key": "value"}).Return(mongo.SingleResult{}, nil)

	// Testing the FindOne method
	result, err := userRepository.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// jsonData, err := json.MarshalIndent(result, "", "    ")
}
