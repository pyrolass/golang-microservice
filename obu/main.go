package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pyrolass/golang-microservice/entities"
)

const sendInterval = 4

const wsEndpoint = "ws://localhost:8080/ws"

func genCoord() float64 {
	n := rand.Intn(100) + 1
	f := rand.Float64()

	return float64(n) + f

}

func genLoc() (float64, float64) {
	return genCoord(), genCoord()
}

func main() {

	for {

		obuIds := generateOBUIds(20)

		conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

		if err != nil {
			log.Fatal("dial:", err)
		}

		for i := 0; i < len(obuIds); i++ {

			lat, long := genLoc()

			data := entities.OBUData{
				OBUID: obuIds[i],
				Lat:   lat,
				Lon:   long,
			}

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("OBU ID: %d, Lat: %f, Lon: %f\n", data.OBUID, data.Lat, data.Lon)
		}
		time.Sleep(sendInterval * time.Second)

	}

}

func generateOBUIds(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())

}
