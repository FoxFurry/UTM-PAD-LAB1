.PHONY: proto-generate build run run-grpc run-docker test bench

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./services/cache/cache.proto

build:
	GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o build/cache
	chmod +x build/cache

run:
	go run main.go serve -p=8080

run-grpc:
	go run main.go serve-grpc

run-docker: build
	docker compose up --build

test:
	go test ./...

bench:
	go test -bench=. ./tests/
