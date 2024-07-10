package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Comment string `from:"text" json:"text"`
}

func main() {

	app := fiber.New()
	api := app.Group("/api/v1")

	api.Post("/comment", createComment)

	app.Listen(":3000")
}

func ConnectProducer(brokerUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerUrl, config)

	return conn, err

}

func PushCommentToKafka(topic string, message []byte) error {
	// code to push comment to kafka
	brokerUrl := []string{"localhost:9092", "localhost:9093"}

	producer, err := ConnectProducer(brokerUrl)

	if err != nil {
		log.Println(err)
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

func createComment(c *fiber.Ctx) error {

	cmt := new(Comment)

	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)

		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

		return err
	}

	cmtInBytes, err := json.Marshal(cmt)

	PushCommentToKafka("comments", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment created successfully",
		"comment": cmt,
	})

	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

		return err
	}

	return err

}
