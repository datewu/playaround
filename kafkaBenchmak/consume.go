package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

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
	// How to decide partition, is it fixed value...?
	//consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	//consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	//consumer, err := master.ConsumePartition(topic, 0, 1844)
	consumer, err := master.ConsumePartition(topic, 0, 1843)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many message processed
	msgCount := 0

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				time.Sleep(time.Second)
				msgCount++
				fmt.Println("Received messages", msg.Offset, string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}
