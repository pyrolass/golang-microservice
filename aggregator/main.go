package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

func main() {

	listenAddr := flag.String("listen-addr", ":8081", "server listen address")

	store := NewMemoryStore()

	aggregator := NewInvoiceAggregator(store)

	makeHttpTransport(*listenAddr, aggregator)

}

func makeHttpTransport(listenAddr string, aggregator AggregatorInterface) {
	logrus.Infof("HTTP transport starting on %s", listenAddr)
	http.HandleFunc("/aggregate", handleAggregation(aggregator))

	http.ListenAndServe(listenAddr, nil)

}

func handleAggregation(aggregator AggregatorInterface) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var distance entities.Distance

		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}

}
