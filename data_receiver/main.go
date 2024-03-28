package main

import (
	"fmt"
	"log"
	"net/http"

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
	prod  DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {

	p, err := NewKafkaProducer(
		kafkaTopic,
	)

	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgch: make(chan entities.OBUData, 128),
		prod:  p,
	}, nil
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

		err = dr.prod.ProduceData(data)

		if err != nil {
			fmt.Printf("kafka produce error: %v\n", err)
		}
	}
}
