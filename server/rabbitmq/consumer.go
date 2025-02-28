package rabbitmq

// package main

import (
	"chat-api/controller"
	"chat-api/model"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func startConsumer(conn *amqp.Connection) {
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	err = declareExchange(ch, "commands_topic")
	if err != nil {
		log.Fatal(err)
	}

	q, err := declareQueue(ch)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		q.Name,           // queue name
		"msg_command",    // routing key
		"commands_topic", // exchange
		false,
		nil)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}

	go func() {
		for data := range msgs {
			msg := model.Message{}
			err := json.Unmarshal(data.Body, &msg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = msg.Create()
			if err != nil {
				fmt.Println(err)
				continue
			}
			controller.ReceivedMsgs <- msg
		}
	}()

	log.Printf("[*] To exit press CTRL+C")
	<-forever
	fmt.Println("chan closed")
}
