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

	// httpClient := client.NewHttpClient(aggregatorURL)
	grpcClient, err := client.NewGRPCClient(aggregatorURL)

	if err != nil {
		logrus.Fatalf("Error creating GRPC Client: %v", err)
	}

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcService, grpcClient)

	if err != nil {
		logrus.Fatalf("Error creating Kafka Consumer: %v", err)
	}

	KafkaConsumer.Start()
}
