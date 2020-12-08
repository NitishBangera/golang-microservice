package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"microservice/src/model"

	"github.com/segmentio/kafka-go"
)

type Queue struct {
	context context.Context
	topic   string
	brokers []string
	groupId string
}

func New(topic string, brokers []string, groupId string) *Queue {
	return &Queue{context: context.Background(), topic: topic, brokers: brokers, groupId: groupId}
}

func (queue *Queue) Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     queue.brokers,
		Topic:       queue.topic,
		GroupID:     queue.groupId,
		StartOffset: kafka.LastOffset,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(queue.context)
		if err != nil {
			panic("Could not read message " + err.Error())
		}
		var eventNotif model.Eventnotification
		if err := json.Unmarshal(msg.Value, &eventNotif); err != nil {
			fmt.Println("failed to unmarshal:", err)
		} else {
			fmt.Println(eventNotif)
		}
	}
}
