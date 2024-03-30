package main

import (
	"context"

	types "github.com/pyrolass/golang-microservice/proto_types"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := types.NewAggregatorClient(conn)

	_, err = c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 10.1,
		Unix:  123456,
	})
}
