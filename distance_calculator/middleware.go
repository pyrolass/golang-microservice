package main

import (
	"time"

	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServiceInterface
}

func NewLogMiddleware(next CalculatorServiceInterface) CalculatorServiceInterface {
	return &LogMiddleware{next: next}
}

func (l *LogMiddleware) CalculateDistance(data entities.OBUData) (distance float64, err error) {

	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{
			"took": time.Since(start).String(),
			"err":  err,
			"dist": distance,
		}).Info("Distance calculated")

	}(time.Now())

	dist, err := l.next.CalculateDistance(data)

	return dist, err
}
