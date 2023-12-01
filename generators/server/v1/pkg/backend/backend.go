package backend

import (
	"context"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

type Backend interface {
	// Create will create resource and returns request.
	Create(ctx context.Context, rs *types.Resource) (*types.Request, error)

	// Delete will delete resource and returns request.
	Delete(ctx context.Context, id *types.Identifier) (*types.Request, error)

	// Update will update resource and returns request.
	Update(ctx context.Context, rs *types.ResourcePatch) (*types.Request, error)

	// Get gets the resource by ID.
	Get(ctx context.Context, id *types.Identifier) (*types.Resource, error)

	// List query resources.
	List(ctx context.Context, query *types.ResourceQuery) (*types.Result, error)

	// GetRequestStatus gets the request status by ID.
	GetRequestStatus(ctx context.Context, requestId string) (*types.Request, error)

	// CancelRequest cancels the request by ID.
	CancelRequest(ctx context.Context, requestId string) (*types.Request, error)

	// ListRequests lists requests.
	ListRequests(ctx context.Context, query *types.RequestQuery) (*types.RequestResult, error)
}

// RequestRepo is the interface to manage the request state
type RequestRepo interface {
	// Save saves the request.
	Save(ctx context.Context, req *types.Request) error

	// GetStatus gets the request status by ID.
	GetStatus(ctx context.Context, requestId string) (*types.Request, error)

	// Cancel cancels the request by ID.
	Cancel(ctx context.Context, requestId string) error

	// List lists requests.
	List(ctx context.Context, query *types.RequestQuery) (*types.RequestResult, error)
}

// ResourceRepo is the interface to manage the resource state
type ResourceRepo interface {
	// Save saves the resource.
	Save(ctx context.Context, resource *types.Resource) error

	// Get gets the resource by ID.
	Get(ctx context.Context, resourceId *types.Identifier) (*types.Resource, error)

	// Delete deletes the resource by ID.
	Delete(ctx context.Context, resourceId *types.Identifier) error

	// List lists resources.
	List(ctx context.Context, query *types.ResourceQuery) (*types.Result, error)
}
