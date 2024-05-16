package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "comments"

	worker, err := connectConsumer([]string{"localhost:9092"})

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	fmt.Println("Consumer started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	maxCount := 0

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)

			case msg := <-consumer.Messages():
				maxCount++
				fmt.Printf("Received message count: %d: | topic (%s) | message: %s\n", maxCount, msg.Topic, string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}

		}
	}()

	<-doneCh
	fmt.Println("Processed", maxCount, "messages")

	if err := worker.Close(); err != nil {
		log.Println(err)
	}
}

func connectConsumer(brokerUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokerUrl, config)

	if err != nil {
		return nil, err
	}

	return consumer, nil
}
