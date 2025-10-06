package request_id

import "context"

const ctxRequestIdKey = "request-id"

func CtxSet(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, ctxRequestIdKey, requestId)
}

func CtxGet(ctx context.Context) string {
	requestId := ctx.Value(ctxRequestIdKey)

	if requestId == nil {
		return ""
	}

	value, ok := requestId.(string)

	if !ok {
		return ""
	}

	return value
}
