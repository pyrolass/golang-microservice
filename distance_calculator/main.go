package main

import "github.com/sirupsen/logrus"

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

const kafkaTopic = "obu-data"

func main() {

	calcService := NewCalculatorService()
	calcService = NewLogMiddleware(calcService)

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService)

	if err != nil {
		logrus.Fatalf("Error creating Kafka Consumer: %v", err)
	}

	KafkaConsumer.Start()
}
