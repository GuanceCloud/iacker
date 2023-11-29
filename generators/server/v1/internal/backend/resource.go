package backend

import (
	"context"
	"fmt"

	errors "github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1/errors"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

func (w *StateBackend) Create(ctx context.Context, rs *types.Resource) (*types.Request, error) {
	ctx, err := w.prepare(ctx)
	if err != nil {
		return nil, err
	}
	return w.apply(ctx, types.NewRequest(rs, types.RequestOperationCreate))
}

func (w *StateBackend) Delete(ctx context.Context, id *types.Identifier) (*types.Request, error) {
	ctx, err := w.prepare(ctx)
	if err != nil {
		return nil, err
	}

	// get resource from graph
	rs, err := w.Resource.Get(ctx, id)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, fmt.Errorf("resource not exists: %v", id)
		}
		return nil, err
	}
	return w.apply(ctx, types.NewRequest(rs, types.RequestOperationDelete))
}

func (w *StateBackend) Update(ctx context.Context, patch *types.ResourcePatch) (*types.Request, error) {
	ctx, err := w.prepare(ctx)
	if err != nil {
		return nil, err
	}

	// get resource from graph
	rs, err := w.Resource.Get(ctx, patch.Id)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, fmt.Errorf("resource not exists: %v", patch.Id)
		}
		return nil, err
	}

	// apply patch to resource
	rs, err = patch.Apply(rs)
	if err != nil {
		return nil, err
	}
	return w.apply(ctx, types.NewRequest(rs, types.RequestOperationUpdate))
}

func (w *StateBackend) Get(ctx context.Context, id *types.Identifier) (*types.Resource, error) {
	ctx, err := w.prepare(ctx)
	if err != nil {
		return nil, err
	}

	return w.Resource.Get(ctx, id)
}

func (w *StateBackend) List(ctx context.Context, query *types.ResourceQuery) (*types.Result, error) {
	ctx, err := w.prepare(ctx)
	if err != nil {
		return nil, err
	}

	return w.Resource.List(ctx, query)
}
