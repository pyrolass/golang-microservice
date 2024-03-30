package client

import (
	types "github.com/pyrolass/golang-microservice/proto_types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	c := types.NewAggregatorClient(conn)

	return &GRPCClient{
		Endpoint:         endpoint,
		AggregatorClient: c,
	}, nil
}
