package backend

import (
	"context"
	"fmt"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

// GetRequestStatus gets the request status by ID.
func (w *StateBackend) GetRequestStatus(ctx context.Context, requestId string) (*types.Request, error) {
	return w.Request.GetStatus(ctx, requestId)
}

// CancelRequest cancels the request by ID.
func (w *StateBackend) CancelRequest(ctx context.Context, requestId string) (*types.Request, error) {
	if err := w.Request.Cancel(ctx, requestId); err != nil {
		return nil, fmt.Errorf("cancel request: %w", err)
	}
	return w.Request.GetStatus(ctx, requestId)
}

// ListRequests lists requests.
func (w *StateBackend) ListRequests(ctx context.Context, query *types.RequestQuery) (*types.RequestResult, error) {
	return w.Request.List(ctx, query)
}
