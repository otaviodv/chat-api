package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitConnection *amqp.Connection

func declareExchange(ch *amqp.Channel, topic string) error {
	return ch.ExchangeDeclare(
		topic,   // name
		"topic", // type
		true,    // durable?
		false,   // auto-deleted?
		false,   // internal?
		false,   // no-wait?
		nil,     // arguments
	)
}

func StartRabbitMQ() {
	var err error
	addr := os.Getenv("RABBITMQ_ADDR")
	rabbitConnection, err = amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s/", addr))
	if err != nil {
		panic(err)
	}

	go startEmitter(rabbitConnection)
}
