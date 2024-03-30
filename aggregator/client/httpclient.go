package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

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

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))

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
