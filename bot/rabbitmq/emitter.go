package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func startEmitter(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	declareExchange(ch, "commands_topic")
	return nil
}

func Push(event []byte) error {
	ch, err := rabbitConnection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	return ch.Publish("commands_topic", "msg_command", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        event,
	})
}
