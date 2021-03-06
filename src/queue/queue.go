package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"microservice/src/model"
	"microservice/src/worker"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

// Queue Structure
type Queue struct {
	context     context.Context
	reader      *kafka.Reader
	writer      *kafka.Writer
	handler     *worker.Handler
	redisClient *redis.Client
}

// NewQueue New method creates a Queue object.
func NewQueue(topic string, brokers []string, groupID string, redisAddress string) *Queue {
	fmt.Println("Initializing listener")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.LastOffset,
	})

	writer := &kafka.Writer{Addr: kafka.TCP(brokers...), Topic: topic}

	handler := worker.NewHandler()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Queue{context: context.Background(), reader: reader, writer: writer, handler: handler, redisClient: redisClient}
}

// Consume method creates a reader from Kafka and reads the messages.
func (queue *Queue) Consume() {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := queue.reader.ReadMessage(queue.context)
		if err != nil {
			panic("Could not read message " + err.Error())
		}
		go func(Value []byte) {
			var eventNotification model.Eventnotification
			if err := json.Unmarshal(Value, &eventNotification); err != nil {
				fmt.Println("Failed to unmarshal:", err)
			} else {
				if eventNotification.EventData == nil {
					val, err := queue.redisClient.Get(queue.context, eventNotification.ReferenceID).Result()
					if err != nil {
						fmt.Println("Cannot read data for reference id from redis :", err, eventNotification)
					} else {
						var eventData map[string]interface{}
						if err := json.Unmarshal([]byte(val), &eventData); err != nil {
							fmt.Println("Failed to unmarshal redis data :", err)
						} else {
							eventNotification.EventData = eventData
						}
					}
				}
				queue.handler.Handle(eventNotification)
			}
		}(msg.Value)
	}
}

// Produce method produces a message to kafka
func (queue *Queue) Produce(eventNotification model.Eventnotification) {
	eventData, err := json.Marshal(eventNotification.EventData)
	if err != nil {
		fmt.Println("Failed to marshal event data :", err)
	} else {
		referenceID := uuid.New().String()
		err := queue.redisClient.Set(queue.context, referenceID, eventData, -1).Err()
		if err != nil {
			fmt.Println("Failed to set event data in redis :", err)
		} else {
			eventNotification.ReferenceID = referenceID
			eventNotification.EventData = nil
		}
	}
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
