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
	
.PHONY: obu aggregator