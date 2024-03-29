package main

import (
	"time"

	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	Next AggregatorInterface
}

func NewLogMiddleware(next AggregatorInterface) AggregatorInterface {
	return &LogMiddleware{
		Next: next,
	}
}

func (l *LogMiddleware) AggregateDistance(data entities.Distance) (err error) {
	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{
			"took": time.Since(start).String(),
			"err":  err,
			"dist": data,
		}).Info("Aggregate Distance")

	}(time.Now())

	err = l.Next.AggregateDistance(data)

	return err

}

func (l *LogMiddleware) CalculateInvoice(obuId int) (inv *entities.Invoice, err error) {
	defer func(start time.Time) {

		var distance float64
		var cost float64

		if inv != nil {
			distance = inv.TotalDistance
			cost = inv.TotalCost
		}

		logrus.WithFields(logrus.Fields{
			"took":           time.Since(start).String(),
			"err":            err,
			"total distance": distance,
			"total cost":     cost,
		}).Info("Distance Sum")

	}(time.Now())

	inv, err = l.Next.CalculateInvoice(obuId)

	return inv, err

}
