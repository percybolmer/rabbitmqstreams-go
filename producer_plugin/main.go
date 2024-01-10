package main

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

const (
	EVENTSTREAM = "events"
)

type Event struct {
	Name string
}

func main() {
	// COnnect to the Stream Plugin on Rabbimq
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))
	if err != nil {
		panic(err)
	}
	// Declare the stream, Set segmentsize and Maxbytes on stream
	err = env.DeclareStream(EVENTSTREAM, stream.NewStreamOptions().
		SetMaxSegmentSizeBytes(stream.ByteCapacity{}.MB(1)).
		SetMaxLengthBytes(stream.ByteCapacity{}.MB(2)))

	if err != nil {
		panic(err)
	}

	// Create a new Producer
	producerOptions := stream.NewProducerOptions()
	producerOptions.SetProducerName("producer")

	// Batch 100 Events in the same Frame, and the SDK will handle everything
	// DEDPULICATION DOES NOT WORK WITH SUBENTRY
	producerOptions.SetSubEntrySize(100)
	producerOptions.SetCompression(stream.Compression{}.Gzip())

	producer, err := env.NewProducer(EVENTSTREAM, producerOptions)
	if err != nil {
		panic(err)
	}

	// Publish 6001 messages
	for i := 0; i <= 6001; i++ {

		event := Event{
			Name: "test",
		}

		data, err := json.Marshal(event)
		if err != nil {
			panic(err)
		}

		message := amqp.NewMessage(data)
		// Apply properties to our message
		props := &amqp.MessageProperties{
			CorrelationID: uuid.NewString(),
		}
		message.Properties = props
		// Set Publishing ID to prevent Deduplication
		message.SetPublishingId(int64(i))
		// Sending the message
		if err := producer.Send(message); err != nil {
			panic(err)
		}
	}

	producer.Close()
}
