package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServiceInterface
}

func NewKafkaConsumer(topic string, cs CalculatorServiceInterface) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"auto.offset.reset": "earliest",
		"group.id":          "myGroup",
	})

	c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{consumer: c, calcService: cs}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Starting Kafka Consumer")
	c.isRunning = true

	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {

	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			logrus.Errorf("kafka consumer error: %v", err)
			continue
		}

		var obuData entities.OBUData

		if err := json.Unmarshal(msg.Value, &obuData); err != nil {
			logrus.Errorf("Error unmarshalling OBU data: %v", err)
			continue
		}

		_, err = c.calcService.CalculateDistance(obuData)

		if err != nil {
			logrus.Errorf("Error calculating distance: %v", err)
			continue
		}

	}
}
