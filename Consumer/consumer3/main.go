package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092,localhost:29092,localhost:39092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic-1"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	i := 1
	for run {
		fmt.Println("Listen on topic-1", i)
		i++
		msg, err := c.ReadMessage(time.Second)
		if err != nil {
			//fmt.Sprintf("%v", err)
			if err.Error() == "Local: Timed out" {
			}
			//fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		} else {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		}

	}

	c.Close()
}
