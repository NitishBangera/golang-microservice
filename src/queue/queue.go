package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"microservice/src/model"

	"github.com/segmentio/kafka-go"
)

// Queue Structure
type Queue struct {
	context context.Context
	reader  *kafka.Reader
	writer  *kafka.Writer
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
	return &Queue{context: context.Background(), reader: reader, writer: writer}
}

// Consume method creates a reader from Kafka and reads the messages.
func (queue *Queue) Consume() {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := queue.reader.ReadMessage(queue.context)
		if err != nil {
			panic("Could not read message " + err.Error())
		}
		var eventNotif model.Eventnotification
		if err := json.Unmarshal(msg.Value, &eventNotif); err != nil {
			fmt.Println("Failed to unmarshal:", err)
		} else {
			fmt.Println("Received :", eventNotif)
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
