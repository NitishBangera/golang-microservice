package main

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-consumer/model"

	"github.com/segmentio/kafka-go"
)

const (
	topic = "search"
)

func getbrokers() []string {
	return []string{"172.16.2.31:9092", "172.16.2.32:9092"}
}

func main() {
	ctx := context.Background()
	consume(ctx)
}

func consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     getbrokers(),
		Topic:       topic,
		GroupID:     "testgroup",
		StartOffset: kafka.LastOffset,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
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
