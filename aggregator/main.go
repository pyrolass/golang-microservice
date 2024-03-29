package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/pyrolass/golang-microservice/entities"
	"github.com/sirupsen/logrus"
)

func main() {

	listenAddr := flag.String("listen-addr", ":3000", "server listen address")

	store := NewMemoryStore()

	aggregator := NewInvoiceAggregator(store)

	aggregator = NewLogMiddleware(aggregator)

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
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			return
		}

		err := aggregator.AggregateDistance(distance)

		if err != nil {
			logrus.Errorf("Error aggregating distance: %s", err)
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			return
		}

	}

}

func writeJson(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
