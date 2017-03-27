package logging

import (
	"context"
	"net/http"
)

type contextKey int

const (
	requestIdKey contextKey = iota + 1
)

func RequestId(r *http.Request) string {
	ctx := r.Context()
	if requestId, ok := ctx.Value(requestIdKey).(string); ok {
		return requestId
	}
	return ""
}

func SetRequestId(r *http.Request, id string) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, requestIdKey, id)
	return r.WithContext(ctx)
}
