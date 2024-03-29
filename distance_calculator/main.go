package main

import "github.com/sirupsen/logrus"

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

const kafkaTopic = "obu-data"

func main() {

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic)

	if err != nil {
		logrus.Fatalf("Error creating Kafka Consumer: %v", err)
	}

	KafkaConsumer.Start()
}
