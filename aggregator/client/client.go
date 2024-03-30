package client

import (
	types "github.com/pyrolass/golang-microservice/proto_types"
	"golang.org/x/net/context"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
