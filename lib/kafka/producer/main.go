package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/nghiant3223/gopractice/kafka/config"
)

func main() {
	configMap := &kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers}
	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Printf("cannot create producer: %v", err)
		return
	}

	for i := 0; i < config.MessageCount; i++ {
		deliveryChannel := make(chan kafka.Event)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &config.Topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf(config.MessageContent, i)),
			Key:   nil,
		}
		err = producer.Produce(msg, deliveryChannel)
		e := <-deliveryChannel
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			return
		}
		log.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
