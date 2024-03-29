package main

import "github.com/pyrolass/golang-microservice/entities"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Insert(data entities.Distance) error {
	return nil
}
