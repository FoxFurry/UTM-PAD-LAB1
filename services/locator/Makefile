.PHONY: proto-generate

proto-generate:
	protoc \
	--go_out=./go-client \
	--go-grpc_out=./go-client \
	ServiceLocator/ServiceLocator/Protos/locator.proto
