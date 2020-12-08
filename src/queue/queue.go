package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"microservice/src/model"
	"microservice/src/worker"

	"github.com/segmentio/kafka-go"
)

// Queue Structure
type Queue struct {
	context context.Context
	reader  *kafka.Reader
	writer  *kafka.Writer
	handler *worker.Handler
}

// New method creates a Queue object.
func New(topic string, brokers []string, groupID string) *Queue {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.LastOffset,
	})

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	handler := worker.New()
	fmt.Println("Initializing listener")
	return &Queue{context: context.Background(), reader: reader, writer: writer, handler: handler}
}

// Consume method creates a reader from Kafka and reads the messages.
func (queue *Queue) Consume() {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := queue.reader.ReadMessage(queue.context)
		if err != nil {
			panic("Could not read message " + err.Error())
		}
		var eventNotification model.Eventnotification
		if err := json.Unmarshal(msg.Value, &eventNotification); err != nil {
			fmt.Println("Failed to unmarshal:", err)
		} else {
			queue.handler.Handle(eventNotification)
		}
	}
}

// Produce method produces a message to kafka
func (queue *Queue) Produce(eventNotification model.Eventnotification) {
	data, err := json.Marshal(eventNotification)
	if err != nil {
		panic("Couldn't marshal data : " + err.Error())
	} else {
		err := queue.writer.WriteMessages(queue.context, kafka.Message{
			Value: data,
		})
		if err != nil {
			panic("Couldn't write message " + err.Error())
		}
	}
}
