package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

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
	rabbitConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	go startEmitter(rabbitConnection)
}
