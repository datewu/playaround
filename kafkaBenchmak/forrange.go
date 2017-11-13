package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

var b int64

func main() {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Retention = time.Second
	config.Consumer.Offsets.CommitInterval = time.Second

	// Specify brokers address. This is default one
	brokers := []string{"localhost:9092"}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "important"
	consumer, err := master.ConsumePartition(topic, 0, 1843)
	if err != nil {
		panic(err)
	}

	// Get signnal for finish
	for msg := range consumer.Messages() {
		time.Sleep(time.Second)
		if msg.Offset == 1846 {
			b = 1846
			if err := consumer.Close(); err != nil {
				panic(err)
			}
			break
		}
		fmt.Println("Received messages id:", msg.Offset, string(msg.Key), string(msg.Value))
	}

	consumer2, err := master.ConsumePartition(topic, 0, b)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer2.Close(); err != nil {
			panic(err)
		}
	}()

	/* can not consumer unter the same master
	go func() {
		consumer3, err := master.ConsumePartition(topic, 0, 1833)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := consumer3.Close(); err != nil {
				panic(err)
			}
		}()
		for msg := range consumer3.Messages() {
			time.Sleep(time.Second)
			fmt.Println("3 id:", msg.Offset, string(msg.Key), string(msg.Value))
		}

	}()
	*/
	fmt.Println("new consumer2")
	for msg := range consumer2.Messages() {
		time.Sleep(2 * time.Second)
		fmt.Println("new Received messages id:", msg.Offset, string(msg.Key), string(msg.Value))
	}
}
