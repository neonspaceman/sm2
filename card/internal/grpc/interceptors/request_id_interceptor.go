package interceptors

import (
	"card/internal/consts"
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"platform/pkg/request_id"
)

func InterceptorRequestId() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		requestId := metadata.ExtractIncoming(ctx).Get(consts.GrpcRequestIdKey)

		if requestId != "" {
			requestId += " "
		}

		requestId += uuid.NewString()

		ctx = request_id.CtxSet(ctx, requestId)

		return handler(ctx, req)
	}
}
