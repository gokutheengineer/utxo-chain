build:
	go build -o bin/utxo-chain

run: build
	./bin/utxo-chain

test:
	go test -v ./...

proto:
	rm -f pb/*.go 
	protoc --proto_path=proto --go_out=./pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

.PHONY: build run test proto