package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

const sendInterval = 4

func genCoord() float64 {
	n := rand.Intn(100) + 1
	f := rand.Float64()

	return float64(n) + f

}

func genLoc() (float64, float64) {
	return genCoord(), genCoord()
}

func main() {

	obuIds := generateOBUIds(20)

	for {

		for i := 0; i < len(obuIds); i++ {

			lat, long := genLoc()

			data := OBUData{
				OBUID: obuIds[i],
				Lat:   lat,
				Lon:   long,
			}

			fmt.Println(data)
		}
		time.Sleep(sendInterval * time.Second)

		fmt.Println(genCoord())

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
