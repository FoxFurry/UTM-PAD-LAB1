.PHONY: proto-generate build run run-docker

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/catalogue.proto

build:
	GOOS=linux GOARCH=amd64 go build -o build/catalogue cmd/catalogue.go

run:
	go run cmd/catalogue.go --config=./config.json

run-docker: build
	docker compose up