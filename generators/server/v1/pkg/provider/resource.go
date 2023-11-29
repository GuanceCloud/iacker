package provider

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/schema"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

// ResourceConfig is the options of resource reconciler.
type ResourceConfig struct {
	// Timeout is the timeout of each operation.
	Timeouts ResourceOperationTimeout

	// Resource is the resource to be reconciled.
	Schema schema.Schema

	// TypeName is the name of resource type.
	TypeName string
}

// ResourceOperationTimeout is the timeout of each operation.
type ResourceOperationTimeout struct {
	// timeout of create operation
	Create time.Duration

	// timeout of update operation
	Update time.Duration

	// timeout of delete operation
	Delete time.Duration

	// timeout of read operation
	Read time.Duration

	// timeout of list operation
	List time.Duration
}

// Factory is the factory of resource provider.
type Factory func(logger log.Logger) (Resource, error)

// Resource define the operations of resource.
type Resource interface {
	// Create creates a resource.
	Create(ctx context.Context, state *State) error

	// Update updates a resource.
	Update(ctx context.Context, state *State) error

	// Delete deletes a resource.
	Delete(ctx context.Context, state *State) error

	// Read reads a resource.
	Read(ctx context.Context, state *State) error

	// List lists resources.
	List(ctx context.Context, query *types.ResourceQuery) (*types.Result, error)

	// Validate validates the resource.
	Validate(ctx context.Context, state *State) error

	// GetConfig returns the options of resource reconciler.
	GetConfig() *ResourceConfig
}
