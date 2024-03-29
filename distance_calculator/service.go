package main

import (
	"math"

	"github.com/pyrolass/golang-microservice/entities"
)

type CalculatorServiceInterface interface {
	CalculateDistance(entities.OBUData) (float64, error)
}

type CalculatorService struct {
	points [][]float64
}

func NewCalculatorService() CalculatorServiceInterface {
	return &CalculatorService{
		points: make([][]float64, 10),
	}
}

func (s *CalculatorService) CalculateDistance(data entities.OBUData) (float64, error) {

	distance := 0.0

	if len(s.points) > 0 {

		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Lat, data.Lon)

	}
	s.points = append(s.points, []float64{data.Lat, data.Lon})
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
