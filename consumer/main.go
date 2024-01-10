package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", "guest", "guest", "localhost:5672", ""))
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	if err := ch.Qos(50, 0, false); err != nil {
		panic(err)
	}

	// Auto ACk has to be FALSE
	stream, err := ch.Consume("events", "events_consumer", false, false, false, false, amqp.Table{
		"x-stream-offset": "10m", // STARTING POINT
	})
	if err != nil {
		panic(err)
	}

	// Loop forever and just read the messages
	fmt.Println("Starting to consume stream")
	for event := range stream {
		fmt.Printf("Event: %s\n", event.CorrelationId)
		fmt.Printf("Headers: %v\n", event.Headers)
		// The payload is in the body
		fmt.Printf("Data: %v\n", string(event.Body))

	}
	ch.Close()
}
