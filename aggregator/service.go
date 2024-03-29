package main

import (
	"fmt"

	"github.com/pyrolass/golang-microservice/entities"
)

type AggregatorInterface interface {
	AggregateDistance(entities.Distance) error
	// DistanceSum(int) (float64, error)
	CalculateInvoice(int) (*entities.Invoice, error)
}

type Storage interface {
	Insert(entities.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storage
}

func NewInvoiceAggregator(store Storage) AggregatorInterface {
	return &InvoiceAggregator{
		store: store,
	}
}

func (s *InvoiceAggregator) AggregateDistance(data entities.Distance) error {
	fmt.Println("Aggregating distance")
	return s.store.Insert(data)

}

func (s *InvoiceAggregator) CalculateInvoice(obuId int) (*entities.Invoice, error) {

	dist, err := s.store.Get(obuId)

	if err != nil {
		return nil, fmt.Errorf("no data found")
	}

	inv := &entities.Invoice{
		OBUID:         obuId,
		TotalDistance: dist,
		TotalCost:     dist * 0.1,
	}

	return inv, nil

}
