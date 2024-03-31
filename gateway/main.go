package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"

	"github.com/pyrolass/golang-microservice/aggregator/client"
	"github.com/sirupsen/logrus"
)

const aggrEndpoint = "http://localhost:3000"

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listen-addr", ":8081", "server listen address")

	flag.Parse()

	client := client.NewHttpClient(aggrEndpoint)

	h := NewInvoiceHandler(client)

	http.HandleFunc("/invoice", makeApiFunc(h.handleGetInvoice))
	logrus.Infof("HTTP transport starting on :8081")
	err := http.ListenAndServe(*listenAddr, nil)

	if err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}

}

type InvoiceHandler struct {
	client client.Client
}

func NewInvoiceHandler(client client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: client,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {

	// access agg client
	invoice, err := h.client.GetInvoice(context.Background(), 1)

	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, invoice)

}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeApiFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			logrus.Errorf("Error handling request: %s", err)
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

}
