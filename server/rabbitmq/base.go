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
	rabbitConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	// go StartConsumer(rabbitConnection)
	go startConsumer(rabbitConnection)

}
