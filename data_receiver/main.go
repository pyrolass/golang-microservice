package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
	"github.com/pyrolass/golang-microservice/entities"
)

const kafkaTopic = "obu-data"

func main() {

	// Produce messages to topic (asynchronously)

	recv, err := NewDataReceiver()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Hello, World!")
}

type DataReceiver struct {
	msgch chan entities.OBUData
	conn  *websocket.Conn
	prod  *kafka.Producer
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	if err != nil {
		return nil, err
	}

	return &DataReceiver{
		msgch: make(chan entities.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data entities.OBUData) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	topic := kafkaTopic

	err = dr.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	if err != nil {
		return err
	}

	return nil

}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {

	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, error := u.Upgrade(w, r, nil)

	if error != nil {
		log.Println(error)
	}

	dr.conn = conn

	go dr.wsReceiveLoop()

}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("client connected")

	for {
		data := entities.OBUData{}
		err := dr.conn.ReadJSON(&data)

		if err != nil {
			log.Println(err)
			continue
		}

		err = dr.produceData(data)

		if err != nil {
			fmt.Printf("kafka produce error: %v\n", err)
		}
	}
}
