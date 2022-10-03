proto-generate:
	protoc \
	--go_opt=paths=source_relative \
	--go_out=. \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/catalogue.proto

run:
	go run cmd/catalogue.go --config=./config.json