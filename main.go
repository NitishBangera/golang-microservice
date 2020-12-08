package main

import "microservice/src/queue"

const (
	topic = "search"
)

func getbrokers() []string {
	return []string{"172.16.2.31:9092", "172.16.2.32:9092"}
}

func main() {
	q := queue.New(topic, getbrokers(), "testgroup")
	q.Consume()
}
