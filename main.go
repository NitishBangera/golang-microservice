package main

import "microservice/src/queue"

const (
	topic     = "search"
	redisAddr = "172.16.2.6:6379"
)

func getbrokers() []string {
	return []string{"172.16.2.31:9092", "172.16.2.32:9092"}
}

func main() {
	q := queue.New(topic, getbrokers(), "testgroup", redisAddr)
	q.Consume()
}
