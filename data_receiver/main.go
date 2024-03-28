package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
	"github.com/pyrolass/golang-microservice/entities"
)

const kafkaTopic = "obu-data"

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

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

	// Produce messages to topic (asynchronously)
	topic := kafkaTopic

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("test producing data"),
	}, nil)

	recv := NewDataReceiver()

	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Hello, World!")
}

type DataReceiver struct {
	msgch chan entities.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan entities.OBUData, 128),
	}
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

		fmt.Printf("OBU ID: %d, Lat: %f, Lon: %f\n", data.OBUID, data.Lat, data.Lon)
		// dr.msgch <- data
	}
}
