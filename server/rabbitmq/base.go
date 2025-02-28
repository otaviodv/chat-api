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

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable?
		false, // auto-deleted?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments
	)
}

func StartRabbitMQ() {
	var err error
	address := os.Getenv("RABBITMQ_ADDR")
	rabbitConnection, err = amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s/", address))
	if err != nil {
		panic(err)
	}

	go startConsumer(rabbitConnection)

}
