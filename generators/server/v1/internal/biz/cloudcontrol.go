package biz

import (
	"context"
	"fmt"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/backend"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

// CloudControlUseCase is a Cloud Control use case.
type CloudControlUseCase struct {
	backend backend.Backend
	log     *log.Helper
}

// NewCloudControlUseCase new a Cloud Control use case.
func NewCloudControlUseCase(backend backend.Backend, logger log.Logger) *CloudControlUseCase {
	return &CloudControlUseCase{backend: backend, log: log.NewHelper(logger)}
}

// CreateResource creates a Resource, and returns the new Request.
func (uc *CloudControlUseCase) CreateResource(ctx context.Context, g *types.Resource) (*types.Request, error) {
	uc.log.WithContext(ctx).Infof("CreateResource: %v bytes", len(g.State))

	// create pending request if resource is not found
	req, err := uc.backend.Create(ctx, g)
	if err != nil {
		return nil, fmt.Errorf("save resource request failed: %w", err)
	}

	return req, nil
}

// GetResource gets a Resource by ID.
func (uc *CloudControlUseCase) GetResource(ctx context.Context, id *types.Identifier) (*types.Resource, error) {
	uc.log.WithContext(ctx).Infof("Get: %v", id)

	// get resource from graph
	rs, err := uc.backend.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

// UpdateResource updates a Resource, and returns the new Request.
func (uc *CloudControlUseCase) UpdateResource(ctx context.Context, g *types.ResourcePatch) (*types.Request, error) {
	uc.log.WithContext(ctx).Infof("UpdateResource: %v", g.Id)

	// create pending request for resource
	req, err := uc.backend.Update(ctx, g)
	if err != nil {
		return nil, fmt.Errorf("save resource request failed: %w", err)
	}
	return req, nil
}

// DeleteResource deletes a Resource, and returns the new Request.
func (uc *CloudControlUseCase) DeleteResource(ctx context.Context, id *types.Identifier) (*types.Request, error) {
	uc.log.WithContext(ctx).Infof("DeleteResource: %v", id)

	// create pending request for resource
	req, err := uc.backend.Delete(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("save resource request failed: %w", err)
	}
	return req, nil
}

// ListResources lists Resources.
func (uc *CloudControlUseCase) ListResources(ctx context.Context, query *types.ResourceQuery) (*types.Result, error) {
	uc.log.WithContext(ctx).Infof("ListResources: %v", query)

	// list from remote
	result, err := uc.backend.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("reconcile failed: %w", err)
	}
	return result, nil
}

// GetRequestStatus gets a Request by ID.
func (uc *CloudControlUseCase) GetRequestStatus(ctx context.Context, id string) (*types.Request, error) {
	uc.log.WithContext(ctx).Infof("GetStatus: %v", id)

	// get request from graph
	req, err := uc.backend.GetRequestStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CancelRequest cancels a Request by ID.
func (uc *CloudControlUseCase) CancelRequest(ctx context.Context, id string) (*types.Request, error) {
	uc.log.WithContext(ctx).Infof("Cancel: %v", id)

	// cancel request from graph
	req, err := uc.backend.CancelRequest(ctx, id)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListRequests lists Requests.
func (uc *CloudControlUseCase) ListRequests(ctx context.Context, query *types.RequestQuery) (*types.RequestResult, error) {
	uc.log.WithContext(ctx).Infof("List: %v", query)

	// list requests from graph
	result, err := uc.backend.ListRequests(ctx, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
