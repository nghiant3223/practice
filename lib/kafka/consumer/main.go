package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nghiant3223/gopractice/kafka/config"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":               config.BootstrapServers,
		"group.id":                        config.Group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"auto.offset.reset":               "earliest",
		"enable.auto.commit":              false,
	}
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Printf("Cannot create consumer: %v", err)
		return
	}

	log.Printf("Create consumer successfully\n")

	err = consumer.SubscribeTopics([]string{config.Topic}, nil)
	if err != nil {
		log.Printf("Cannot subscribe to topic: %v", err)
		return
	}

	log.Printf("Subscrube to topic successfully")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run == true {
		select {
		case sig := <-signalChannel:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		case ev := <-consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				log.Printf("Partitions assigned: %v\n", e)
				consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				log.Printf("Partitions revoked: %v\n", e)
				consumer.Unassign()
			case *kafka.Message:
				log.Printf("Message on %s @ %d:\n%s\n", e.TopicPartition, e.TopicPartition.Offset, string(e.Value))
				processMessage(string(e.Value))
				consumer.CommitMessage(e)
			case kafka.PartitionEOF:
				log.Printf("Partition EOF reached: %v\n", e)
			case kafka.Error:
				log.Printf("Error: %v\n", e)
			}
		}
	}

	log.Printf("Closing consumer\n")
	err = consumer.Close()
	if err != nil {
		log.Printf("Cannot close consumer: %v\n", err)
		return
	}
	log.Printf("Consumer closed\n")
}

func processMessage(msg string) {
	log.Printf("Start processing message %s\n", msg)
	time.Sleep(config.ProcessingDuration)
	log.Printf("Processing message %s\n successfully", msg)
}
