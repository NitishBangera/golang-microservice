package worker

import (
	"fmt"
	"microservice/src/model"
)

// Worker interface
type Worker interface {
	Work(eventNotification model.Eventnotification) (model.Eventnotification, error)
	GetEventType() string
}

// TestWorker for work
type TestWorker struct {
	eventType string `default:"test"`
}

// GetEventType of TestWorker
func (worker *TestWorker) GetEventType() string {
	return worker.eventType
}

// Work of TestWorker
func (worker *TestWorker) Work(eventNotification model.Eventnotification) (model.Eventnotification, error) {
	fmt.Println("Processing ", eventNotification.Ttid, eventNotification.Type)
	return model.Eventnotification{Ttid: eventNotification.Ttid}, nil
}
