package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer  *kafka.Consumer
	isRunning bool
}

func NewKafkaConsumer(topic string) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"auto.offset.reset": "earliest",
		"group.id":          "myGroup",
	})

	c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{consumer: c}, nil
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

		logrus.Infof("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

	}
}
