obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
	@go build -o bin/data_receiver ./data_receiver
	@./bin/data_receiver

calculator:
	@go build -o bin/distance_calculator ./distance_calculator
	@./bin/distance_calculator

agg:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto_types/ptypes.proto
	
.PHONY: obu aggregator