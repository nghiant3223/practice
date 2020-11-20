package config

import "time"

var (
	BootstrapServers   = "localhost:9092"
	Topic              = "locations"
	ReplicationFactor  = 1
	PartitionCount     = 3
	MessageCount       = 1000
	MessageContent     = "message %v"
	Group              = "consumer"
	ProcessingDuration = 100 * time.Millisecond
)
