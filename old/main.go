package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	println(uri)

	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	//if err != nil {
	//	panic(err)
	//}
	//
	//coll := client.Database("sample_mflix").Collection("movies")
	//
	//println(coll)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
