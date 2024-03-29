package main

import (
	"fmt"

	"github.com/pyrolass/golang-microservice/entities"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (ms *MemoryStore) Insert(data entities.Distance) error {
	ms.data[data.OBUID] += data.Value
	return nil
}

func (ms *MemoryStore) Get(id int) (float64, error) {
	distance, ok := ms.data[id]
	if !ok {
		return 0, fmt.Errorf("no data found")
	}

	return distance, nil
}
