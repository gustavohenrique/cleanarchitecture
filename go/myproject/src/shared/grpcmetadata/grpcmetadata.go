package grpcmetadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func GetRolesFrom(ctx context.Context) string {
	return GetValue(ctx, "roles")
}

func GetUserIdFrom(ctx context.Context) string {
	return GetValue(ctx, "user_id")
}

func GetValue(ctx context.Context, key string) string {
	md, _ := metadata.FromIncomingContext(ctx)
	value := md.Get(key)
	if len(value) > 0 {
		return value[0]
	}
	return ""
}
