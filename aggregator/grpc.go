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

func (s *GRPCServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := entities.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.scv.AggregateDistance(distance)
}
