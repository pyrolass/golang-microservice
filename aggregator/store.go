package main

import (
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

func (s *MemoryStore) Insert(data entities.Distance) error {
	s.data[data.OBUID] += data.Value
	return nil
}
