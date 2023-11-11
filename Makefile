build:
	go build -o bin/utxo-chain

run: build
	./bin/utxo-chain

test:
	go test ./...

PHONY: build run test