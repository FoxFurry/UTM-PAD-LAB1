.PHONY: proto-generate run-db run-local

proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./services/user/user.proto

run:
	go run cmd/user.go

run-db:
	docker compose up --build -d db

run-local: run-db run
