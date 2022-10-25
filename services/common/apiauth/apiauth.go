package apiauth

// REQ: Cache auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ValidateContextKey(ctx context.Context, expectedKey string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("failed to parse cache metadata")
	}

	return ValidateMetadataKey(md, expectedKey)
}

func ValidateMetadataKey(md metadata.MD, expectedKey string) error {
	keyMetadata := md["key"]
	if len(keyMetadata) == 0 {
		return fmt.Errorf("key metadata is missing")
	}

	if md["key"][0] != expectedKey {
		return fmt.Errorf("key mismatch")
	}

	return nil
}

func SetRequestMetadataKey(ctx context.Context, key string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"key": key}))
}

func SetResponseMetadataKey(ctx context.Context, key string) (context.Context, error) {
	if err := grpc.SendHeader(ctx, metadata.New(map[string]string{"key": key})); err != nil {
		return nil, fmt.Errorf("failed to send response metadata: %w", err)
	}

	return ctx, nil
}
