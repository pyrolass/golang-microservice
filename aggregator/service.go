package main

import (
	"fmt"

	"github.com/pyrolass/golang-microservice/entities"
)

type AggregatorInterface interface {
	AggregateDistance(entities.Distance) error
}

type Storage interface {
	Insert(entities.Distance) error
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
