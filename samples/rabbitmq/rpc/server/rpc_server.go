package main

import (
	"context"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(1, 0, false)

	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))

			failOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)

			response := fib(n)

			err = ch.PublishWithContext(ctx, "", d.ReplyTo,
				false, false, amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}

	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever

}
