package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pyrolass/golang-microservice/entities"
)

func main() {

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
		msgch: make(chan entities.OBUData),
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
		dr.msgch <- data
	}
}
