package main

import (
	"fmt"
	"microservice/src/model"
	"microservice/src/queue"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var config *model.Config

const (
	topic            = "search"
	mqPort           = "9092"
	redisPort        = "6379"
	redisClusterPort = "6381"
	groupID          = "testgroup"
)

func main() {
	cfg, err := ini.Load("/etc/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	values := make(map[string]string)
	for _, val := range cfg.Section("").Keys() {
		value := strings.Replace(strings.Replace(val.Value(), "[", "", -1), "]", "", -1)
		values[val.Name()] = value
	}
	config = model.New(values)

	mqNodes := strings.Split(config.GetValue("MQ_NODES"), ",")
	brokers := make([]string, len(mqNodes))
	for i, val := range mqNodes {
		brokers[i] = val + ":" + mqPort
	}

	redisNode := config.GetValue("REDIS_NODES") + ":" + redisPort

	q := queue.New(topic, brokers, groupID, redisNode)
	q.Consume()
}
