package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}

	}()

	//Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1: // rabbitmq
			fmt.Printf("received from RabbitMq: %s ID: %d\n", msg1.Msg, msg1.id)
		case msg2 := <-c2: // kafka
			fmt.Printf("received from Kafka: %s ID: %d\n", msg2.Msg, msg2.id)
		case <-time.After(time.Second * 3):
			println("tinmeout")
		}
	}

}
