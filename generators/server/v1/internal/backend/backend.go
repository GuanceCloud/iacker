package backend

import (
	"context"
	"fmt"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/backend"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/provider"
)

type StateBackend struct {
	Request     backend.RequestRepo
	Resource    backend.ResourceRepo
	Resources   map[string]provider.Resource
	Middlewares []provider.Middleware
}

func (w *StateBackend) prepare(ctx context.Context) (context.Context, error) {
	var err error
	for i, middleware := range w.Middlewares {
		ctx, err = middleware.Prepare(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to prepare with middleware %d: %w", i, err)
		}
	}
	return ctx, nil
}
