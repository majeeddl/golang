package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Body      string             `bson:"body"`
	Tags      []string           `bson:"tags"`
	Comments  uint64             `bson:"comments"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	collection := client.Database("test").Collection("trainers")

	_, err = collection.InsertOne(context.TODO(), Trainer{
		ID:        primitive.NewObjectID(),
		Title:     "Title 1",
		Body:      "Body 1",
		Tags:      []string{"tag1", "tag2"},
		Comments:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err)
	}

	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{
		{"title", "Title 1"},
	}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Fatal("No document found")
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonData))

}
