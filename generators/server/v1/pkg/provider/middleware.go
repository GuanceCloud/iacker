package provider

import (
	"context"
)

// Middleware is the interface to build the context
type Middleware interface {
	// Prepare is the method to implement the prepare logic
	// Such as authentication
	// It will be called before the request is persisted
	// All the result should be bound to the context
	Prepare(ctx context.Context) (context.Context, error)
}

// BaseMiddleware is the base struct for middleware developing
type BaseMiddleware struct{}

// Prepare is the method to implement the prepare logic
func (b BaseMiddleware) Prepare(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
