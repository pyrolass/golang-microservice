package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pyrolass/golang-microservice/entities"
)

type DataProducer interface {
	ProduceData(entities.OBUData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(topic string) (DataProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					// fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer: p, topic: topic}, nil
}

func (kp *KafkaProducer) ProduceData(data entities.OBUData) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
