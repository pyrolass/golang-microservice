package main

import (
	"encoding/json"
	"flag"
	"net"
	"net/http"
	"strconv"

	"github.com/pyrolass/golang-microservice/entities"
	types "github.com/pyrolass/golang-microservice/proto_types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	listenAddr := flag.String("listen-addr", ":3000", "server listen address")
	grpcListenAddr := flag.String("grpc-addr", ":3001", "server listen address")

	store := NewMemoryStore()

	aggregator := NewInvoiceAggregator(store)

	aggregator = NewLogMiddleware(aggregator)

	go func() {
		err := makeGRPCTransport(*grpcListenAddr, aggregator)

		if err != nil {
			logrus.Fatalf("Error creating GRPC transport: %v", err)
		}

	}()

	makeHttpTransport(*listenAddr, aggregator)

}

func makeGRPCTransport(listenAddr string, aggregator AggregatorInterface) error {
	// make the tcp
	logrus.Infof("GRPC transport starting on %s", listenAddr)

	ln, err := net.Listen("tcp", listenAddr)

	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
		return err
	}

	defer ln.Close()
	// make grpc server
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	// register the server
	types.RegisterAggregatorServer(grpcServer, NewGRPCAggregatorServer(aggregator))

	return grpcServer.Serve(ln)

}

func makeHttpTransport(listenAddr string, aggregator AggregatorInterface) {
	logrus.Infof("HTTP transport starting on %s", listenAddr)
	http.HandleFunc("/aggregate", handleAggregation(aggregator))
	http.HandleFunc("/invoice", handleGetInvoice(aggregator))

	http.ListenAndServe(listenAddr, nil)

}

func handleGetInvoice(aggregator AggregatorInterface) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		obuIdValue := r.URL.Query().Get("obuId")
		obuId, _ := strconv.Atoi(obuIdValue)

		invoice, err := aggregator.CalculateInvoice(obuId)

		if err != nil {
			logrus.Errorf("Error getting distance sum: %s", err)
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			return
		}

		writeJson(w, http.StatusOK, map[string]any{"Invoice": invoice})

	}

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
