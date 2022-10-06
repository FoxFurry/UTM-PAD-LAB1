.PHONY: proto-generate build run run-db run-docker run-local

proto-generate:
	protoc \
	-I . \
	-I ../common/proto/common \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/catalogue/catalogue.proto

build:
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o build/catalogue cmd/catalogue.go
	chmod +x build/catalogue

run:
	go run cmd/catalogue.go

run-db:
	docker compose up --build -d db

run-docker: build
	docker compose up --build

run-local: run-db run
