.PHONY: proto-generate build run run-docker

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./services/cache/cache.proto

build:
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o build/cache cmd/cache.go
	chmod +x build/catalogue

run:
	go run cmd/cache.go

run-docker: build
	docker compose up --build
