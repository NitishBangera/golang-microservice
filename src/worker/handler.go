package worker

import (
	"fmt"
	"microservice/src/model"
)

// Handler of workers
type Handler struct {
	workers map[string]Worker
}

// NewHandler New method creates a Handler object.
func NewHandler() *Handler {
	workers := make(map[string]Worker)
	testWorker := new(TestWorker)
	workers[testWorker.GetEventType()] = testWorker
	fmt.Println("Worker :", testWorker.GetEventType())
	return &Handler{workers: workers}
}

// Handle the message.
func (handler *Handler) Handle(eventNotification model.Eventnotification) {
	worker, workerExists := handler.workers[eventNotification.Type]
	if workerExists {
		worker.Work(eventNotification)
	}
}
