package main

import (
	"github.com/pyrolass/golang-microservice/aggregator/client"
	"github.com/sirupsen/logrus"
)

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

const kafkaTopic = "obu-data"
const aggregatorURL = "http://localhost:3000/aggregate"

func main() {

	calcService := NewCalculatorService()
	calcService = NewLogMiddleware(calcService)

	client := client.NewClient(aggregatorURL)

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService, client)

	if err != nil {
		logrus.Fatalf("Error creating Kafka Consumer: %v", err)
	}

	KafkaConsumer.Start()
}
