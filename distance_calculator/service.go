package main

import (
	"math"

	"github.com/pyrolass/golang-microservice/entities"
)

type CalculatorServiceInterface interface {
	CalculateDistance(entities.OBUData) (float64, error)
}

type CalculatorService struct {
	prevPoint []float64
}

func NewCalculatorService() CalculatorServiceInterface {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data entities.OBUData) (float64, error) {

	distance := 0.0

	if len(s.prevPoint) > 0 {

		distance = calculateDistance(s.prevPoint[0], s.prevPoint[0], data.Lat, data.Lon)

	}
	s.prevPoint = []float64{data.Lat, data.Lon}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
