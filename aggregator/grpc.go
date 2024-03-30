package main

import (
	"context"

	"github.com/pyrolass/golang-microservice/entities"
	types "github.com/pyrolass/golang-microservice/proto_types"
)

type GRPCServer struct {
	types.UnimplementedAggregatorServer
	scv AggregatorInterface
}

func NewGRPCAggregatorServer(svc AggregatorInterface) *GRPCServer {
	return &GRPCServer{scv: svc}
}

func (s *GRPCServer) GRPCAggregateDistance(ctx context.Context, req *types.AggregateRequest) error {
	distance := entities.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.scv.AggregateDistance(distance)
}
