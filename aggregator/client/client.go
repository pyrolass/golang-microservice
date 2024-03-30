package client

import (
	"github.com/pyrolass/golang-microservice/entities"
	types "github.com/pyrolass/golang-microservice/proto_types"
	"golang.org/x/net/context"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
	GetInvoice(context.Context, int) (*entities.Invoice, error)
}
