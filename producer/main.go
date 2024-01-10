package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Event struct {
	Name string
}

func main() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", "guest", "guest", "localhost:5672", ""))
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// Create a new Queue
	q, err := ch.QueueDeclare("events", true, false, false, true, amqp.Table{
		"x-queue-type":                    "stream",
		"x-stream-max-segment-size-bytes": 30000,  // EACH SEGMENT FILE IS ALLOWED 0.03 MB
		"x-max-length-bytes":              150000, // TOTAL STREAM SIZE IS 0.15 MB
	})
	if err != nil {
		panic(err)
	}

	// Publish 1001 messages
	ctx := context.Background()
	for i := 0; i <= 1000; i++ {

		event := Event{
			Name: "test",
		}

		data, err := json.Marshal(event)
		if err != nil {
			panic(err)
		}

		err = ch.PublishWithContext(ctx, "", "events", false, false, amqp.Publishing{
			Body:          data,
			CorrelationId: uuid.NewString(),
		})
		if err != nil {
			panic(err)
		}
	}
	// Close the channel to await all messages being sent
	ch.Close()
	fmt.Println(q.Name)
}
