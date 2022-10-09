.PHONY: proto-generate build run run-db run-docker run-local

build:
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o build/catalogue cmd/catalogue.go
	chmod +x build/catalogue

run:
	go run cmd/gateway.go
