package main

import (
	"time"

	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (l *LogMiddleware) ProduceData(data entities.OBUData) error {
	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{
			"obu_id": data.OBUID,
			"lat":    data.Lat,
			"lon":    data.Lon,
			"took":   time.Since(start).String(),
		}).Info("Producing data")

	}(time.Now())

	return l.next.ProduceData(data)
}
