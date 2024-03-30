package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pyrolass/golang-microservice/entities"
	types "github.com/pyrolass/golang-microservice/proto_types"
)

type HttpClient struct {
	Endpoint string
}

func NewHttpClient(endpoint string) *HttpClient {
	return &HttpClient{
		Endpoint: endpoint,
	}
}

func (c *HttpClient) Aggregate(ctx context.Context, data *types.AggregateRequest) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint+"/aggregate", bytes.NewReader(b))

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	// resp, err := httpc.Do(req)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil

}

func (c *HttpClient) GetInvoice(ctx context.Context, id int) (*entities.Invoice, error) {

	invReq := types.GetInvoicerequest{
		ObuID: int32(id),
	}
	b, err := json.Marshal(&invReq)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Endpoint+"/invoice", bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// resp, err := httpc.Do(req)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var inv entities.Invoice

	if err := json.NewDecoder(resp.Body).Decode(&inv); err != nil {
		return nil, err
	}

	return &inv, nil

}
