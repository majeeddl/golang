package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func TestMongoRepository(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// Mocking the mongo client and collection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	assert.Nil(t, err)

	// Creating the repository
	userRepository := MongoRepository[User](client, "test", "users")

	// Mocking the collection's InsertOne method
	// collection.On("InsertOne", ctx, bson.M{"key": "value"}).Return(nil, nil)

	// Testing the InsertOne method
	_, err = userRepository.InsertOne(ctx, User{ID: primitive.NewObjectID(), Name: "masoumeh"})

	assert.Nil(t, err)

	// Mocking the collection's FindOne method
	// userRepository.On("FindOne", ctx, bson.M{"key": "value"}).Return(mongo.SingleResult{}, nil)

	// Testing the FindOne method
	result, err := userRepository.Find(ctx, bson.M{"name": "majeed"})

	fmt.Println(result)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	jsonData, err := json.MarshalIndent(result, "", "    ")

	// data := string(jsonData)

	assert.Nil(t, err)
	assert.NotNil(t, jsonData)

	assert.Nil(t, err)
}

// Mock testing for mongo repository
// func TestFindOne(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("success", func(mt *mtest.T) {
// 		userCollection = mt.Coll
// 		expectedUser := user{
// 			ID:    primitive.NewObjectID(),
// 			Name:  "john",
// 			Email: "john.doe@test.com",
// 		}

// 		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
// 			{"_id", expectedUser.ID},
// 			{"name", expectedUser.Name},
// 			{"email", expectedUser.Email},
// 		}))
// 		userResponse, err := getFromID(expectedUser.ID)
// 		assert.Nil(t, err)
// 		assert.Equal(t, &expectedUser, userResponse)
// 	})
// }
