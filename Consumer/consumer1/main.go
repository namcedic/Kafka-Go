package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type User struct {
	Id       int
	FullName string
	Email    string
	Address  string
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:19092,localhost:29092,localhost:39092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	topic := "topic-1"
	c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		fmt.Println("Listening.....")
		msg, err := c.ReadMessage(time.Second)
		if err != nil {
			fmt.Sprintf("%v", err)
			if err.Error() == "Local: Timed out" {
				//fmt.Println("waiting for message")
			}
			//fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		} else {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		}

	}

	c.Close()
}
