package worker

import (
	"fmt"
	"microservice/src/model"
)

// Handler of workers
type Handler struct {
	workers map[string]interface{}
}

// New method creates a Handler object.
func New() *Handler {
	workers := make(map[string]interface{})
	testWorker := TestWorker{}
	workers[testWorker.GetEventType()] = testWorker
	return &Handler{workers: workers}
}

// Handle the message.
func (handler *Handler) Handle(eventNotification model.Eventnotification) {
	worker, workerExists := handler.workers[eventNotification.Type]
	fmt.Println("Ttid : ", eventNotification.Ttid, "Type : ", eventNotification.Type, " Worker Exists : ", workerExists)
	if workerExists {
		obj, methodExists := worker.(interface {
			Work(eventNotification model.Eventnotification) (model.Eventnotification, error)
		})
		if methodExists {
			obj.Work(eventNotification)
		}
	}
}
