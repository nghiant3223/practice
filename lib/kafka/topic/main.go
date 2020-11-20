package main

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nghiant3223/gopractice/kafka/config"
)

func main() {
	configMap := &kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers}
	admin, err := kafka.NewAdminClient(configMap)
	if err != nil {
		log.Printf("cannot create admin client: %v\n", err)
		return
	}
	defer admin.Close()

	ctx := context.Background()
	topicSpec := kafka.TopicSpecification{
		Topic:             config.Topic,
		NumPartitions:     config.PartitionCount,
		ReplicationFactor: config.ReplicationFactor,
	}
	results, err := admin.CreateTopics(ctx, []kafka.TopicSpecification{topicSpec})
	if err != nil {
		log.Printf("can't create topic: %v\n", err)
		return
	}
	log.Printf("create topic result: %v\n", results[0])
}
