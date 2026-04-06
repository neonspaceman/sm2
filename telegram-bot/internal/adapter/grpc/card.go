package grpc

import (
	"card/pkg/api/card"
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"platform/pkg/request_id"
	card_client "telegram-bot/internal/client/card"
	"telegram-bot/internal/config"
	"telegram-bot/internal/consts"
)

type CardClient struct {
	grpc card.CardServiceClient
}

func NewCardClient(cfg *config.Config) *CardClient {
	conn, err := grpc.NewClient(
		cfg.GRPC.CardURI,
		//grpc.WithChainUnaryInterceptor(InterceptorRequestId()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		panic(err)
	}

	grpcClient := card.NewCardServiceClient(conn)

	return &CardClient{grpc: grpcClient}
}

func (c *CardClient) Create(ctx context.Context, userId uuid.UUID, question, answer string) (*card_client.Card, error) {
	req := &card.CreateRequest{
		UserId:   userId.String(),
		Question: question,
		Answer:   answer,
	}

	response, err := c.grpc.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	return card_client.ToCard(response.Card), nil
}

func (c *CardClient) GetCards(ctx context.Context, userId uuid.UUID, limit uint64, after string) ([]*card_client.Card, bool, string, error) {
	req := &card.GetByUserIdRequest{
		UserId: userId.String(),
		Limit:  limit,
		After:  after,
	}

	response, err := c.grpc.GetByUserId(ctx, req)

	if err != nil {
		return nil, false, "", err
	}

	cards := make([]*card_client.Card, 0, len(response.Cards))

	for _, v := range response.Cards {
		cards = append(cards, card_client.ToCard(v))
	}

	return cards, response.HasNext, response.EndCursor, nil
}

func InterceptorRequestId() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		requestId := request_id.CtxGet(ctx)

		if requestId != "" {
			requestId += ","
		}

		requestId += uuid.NewString()

		ctx = metadata.ExtractOutgoing(ctx).Set(consts.GrpcRequestIdKey, requestId).ToOutgoing(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
