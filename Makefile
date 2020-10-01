.PHONY: clean build

clean:
	rm -rf ./bin/*

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/services/api ./services/api
	GOOS=darwin GOARCH=amd64 go build -o bin/services/processor ./services/processor
	GOOS=darwin GOARCH=amd64 go build -o bin/services/aggregator ./services/aggregator

test:
	go test ./...

run-api:
	./bin/services/api

run-processor:
	./bin/services/processor

run-aggregator:
	./bin/services/aggregator

build-api-docker:
	docker build . -t co2-api -f Dockerfile-Api

build-processor-docker:
	docker build . -t co2-processor -f Dockerfile-Processor

build-aggregator-docker:
	docker build . -t co2-aggregator -f Dockerfile-Aggregator
