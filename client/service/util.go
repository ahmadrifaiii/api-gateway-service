package service

import (
	"context"
	"erp-gateway-service/internal/pkg/app"

	"google.golang.org/grpc/metadata"
)

func setMetadata(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{
		"token": ctx.Value(app.Ctx("token")).(string),
	})

	return metadata.NewOutgoingContext(ctx, md)
}
