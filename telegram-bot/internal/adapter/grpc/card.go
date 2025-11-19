package grpc

import (
	"card/pkg/api/card"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	cardPkg "telegram-bot/internal/client/card"
)

type CardClient struct {
	grpc card.CardServiceClient
}

func NewCardClient() *CardClient {
	conn, err := grpc.NewClient(
		"card:50051",
		//grpc.WithChainUnaryInterceptor(InterceptorRequestId()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		panic(err)
	}

	grpcClient := card.NewCardServiceClient(conn)

	return &CardClient{grpc: grpcClient}
}

func (c *CardClient) Create(ctx context.Context, req cardPkg.CreateRequestDto) error {
	grpcRequest := &card.CreateRequest{
		UserId:   req.UserId,
		Question: req.Question,
		Answer:   req.Answer,
	}

	switch req.FileType {
	case cardPkg.FileTypePhoto:
		grpcRequest.FileId = req.FileId
		grpcRequest.FileType = card.FileType_PHOTO
	case cardPkg.FileTypeDocument:
		grpcRequest.FileId = req.FileId
		grpcRequest.FileType = card.FileType_DOCUMENT
	}

	_, err := c.grpc.Create(ctx, grpcRequest)

	if err != nil {
		return err
	}

	return nil
}

//func InterceptorRequestId() grpc.UnaryClientInterceptor {
//	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//		requestId := request_id.CtxGet(ctx)
//
//		if requestId != "" {
//			requestId += ","
//		}
//
//		requestId += uuid.NewString()
//
//		ctx = metadata.ExtractOutgoing(ctx).Set(consts.GrpcRequestIdKey, requestId).ToOutgoing(ctx)
//
//		return invoker(ctx, method, req, reply, cc, opts...)
//	}
//}
