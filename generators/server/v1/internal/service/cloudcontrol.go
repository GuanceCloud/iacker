package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	ikctx "github.com/GuanceCloud/iacker/generators/server/v1/pkg/provider"

	v1 "github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1"
	"github.com/GuanceCloud/iacker/generators/server/v1/internal/biz"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

// CloudControlService is a service to implement Cloud Control API.
type CloudControlService struct {
	v1.UnimplementedCloudControlServiceServer

	uc *biz.CloudControlUseCase
}

// NewCloudControlService new a greeter service.
func NewCloudControlService(uc *biz.CloudControlUseCase) *CloudControlService {
	return &CloudControlService{uc: uc}
}

// CreateResource creates a new resource.
func (svc *CloudControlService) CreateResource(ctx context.Context, req *v1.CreateResourceRequest) (*v1.CreateResourceResponse, error) {
	resourceRequest, err := svc.uc.CreateResource(
		ikctx.WithContext(ctx, &ikctx.Context{
			Region:      req.Region,
			TypeName:    req.TypeName,
			AccessToken: req.AccessToken,
		}),
		&types.Resource{
			Identifier: &types.Identifier{
				Partition:    "-",
				Service:      "-",
				Region:       req.Region,
				Owner:        "-",
				ResourceType: req.TypeName,
				ResourceId:   "-",
			},
			State:     req.DesiredState,
			CreatedAt: time.Now(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create resource failed: %w", err)
	}

	return &v1.CreateResourceResponse{
		ProgressEvent: buildProgressEvent(resourceRequest),
	}, nil
}

// GetResource gets the resource by name.
func (svc *CloudControlService) GetResource(ctx context.Context, req *v1.GetResourceRequest) (*v1.GetResourceResponse, error) {
	id, err := types.ParseIdentifier(req.Identifier)
	if err != nil {
		return nil, fmt.Errorf("parse identifier failed: %w", err)
	}

	resource, err := svc.uc.GetResource(
		ikctx.WithContext(ctx, &ikctx.Context{
			Region:      id.Region,
			TypeName:    id.ResourceType,
			AccessToken: req.AccessToken,
		}),
		&id,
	)
	if err != nil {
		return nil, fmt.Errorf("get resource failed: %w", err)
	}

	return &v1.GetResourceResponse{
		ResourceDescription: buildResourceDescription(resource),
		TypeName:            resource.Identifier.ResourceType,
	}, nil
}

// DeleteResource deletes the resource by name.
func (svc *CloudControlService) DeleteResource(ctx context.Context, req *v1.DeleteResourceRequest) (*v1.DeleteResourceResponse, error) {
	id, err := types.ParseIdentifier(req.Identifier)
	if err != nil {
		return nil, fmt.Errorf("parse identifier failed: %w", err)
	}

	resourceRequest, err := svc.uc.DeleteResource(
		ikctx.WithContext(ctx, &ikctx.Context{
			Region:      id.Region,
			TypeName:    id.ResourceType,
			AccessToken: req.AccessToken,
		}),
		&id,
	)
	if err != nil {
		return nil, fmt.Errorf("delete resource failed: %w", err)
	}

	return &v1.DeleteResourceResponse{
		ProgressEvent: buildProgressEvent(resourceRequest),
	}, nil
}

// ListResources lists all resources.
func (svc *CloudControlService) ListResources(ctx context.Context, req *v1.ListResourcesRequest) (*v1.ListResourcesResponse, error) {
	if req.MaxResults == 0 {
		req.MaxResults = 100 // default max results
	}

	result, err := svc.uc.ListResources(
		ikctx.WithContext(ctx, &ikctx.Context{
			Region:      req.Region,
			TypeName:    req.TypeName,
			AccessToken: req.AccessToken,
		}),
		&types.ResourceQuery{TypeName: req.TypeName, MaxResults: int(req.MaxResults)},
	)
	if err != nil {
		return nil, err
	}

	resourceDescriptions := make([]*v1.ResourceDescription, 0)
	for _, resource := range result.Items {
		resourceDescriptions = append(resourceDescriptions, buildResourceDescription(resource))
	}

	return &v1.ListResourcesResponse{
		ResourceDescriptions: resourceDescriptions,
		TypeName:             req.TypeName,
		NextToken:            result.NextToken,
	}, nil
}

// UpdateResource updates the resource.
func (svc *CloudControlService) UpdateResource(ctx context.Context, req *v1.UpdateResourceRequest) (*v1.UpdateResourceResponse, error) {
	id, err := types.ParseIdentifier(req.Identifier)
	if err != nil {
		return nil, fmt.Errorf("parse identifier failed: %w", err)
	}

	resourceRequest, err := svc.uc.UpdateResource(
		ikctx.WithContext(ctx, &ikctx.Context{
			Region:      id.Region,
			TypeName:    id.ResourceType,
			AccessToken: req.AccessToken,
		}),
		&types.ResourcePatch{
			Id:      &id,
			Patches: json.RawMessage(req.PatchDocument),
		},
	)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateResourceResponse{
		ProgressEvent: buildProgressEvent(resourceRequest),
	}, nil
}

// CancelResourceRequest cancels the specified resource request.
func (svc *CloudControlService) CancelResourceRequest(ctx context.Context, req *v1.CancelResourceRequestRequest) (*v1.CancelResourceRequestResponse, error) {
	resourceRequest, err := svc.uc.CancelRequest(
		ctx,
		req.RequestToken,
	)
	if err != nil {
		return nil, err
	}
	return &v1.CancelResourceRequestResponse{
		ProgressEvent: buildProgressEvent(resourceRequest),
	}, nil
}

// GetResourceRequestStatus gets the specified resource request status.
func (svc *CloudControlService) GetResourceRequestStatus(ctx context.Context, req *v1.GetResourceRequestStatusRequest) (*v1.GetResourceRequestStatusResponse, error) {
	resourceRequest, err := svc.uc.GetRequestStatus(
		ctx,
		req.RequestToken,
	)
	if err != nil {
		return nil, err
	}
	return &v1.GetResourceRequestStatusResponse{
		ProgressEvent: buildProgressEvent(resourceRequest),
	}, nil
}

// ListResourceRequests list resource requests.
func (svc *CloudControlService) ListResourceRequests(ctx context.Context, req *v1.ListResourceRequestsRequest) (*v1.ListResourceRequestsResponse, error) {
	listResult, err := svc.uc.ListRequests(
		ctx,
		&types.RequestQuery{},
	)
	if err != nil {
		return nil, err
	}

	progressEvents := make([]*v1.ProgressEvent, 0)
	for _, resourceRequest := range listResult.Items {
		progressEvents = append(progressEvents, buildProgressEvent(resourceRequest))
	}
	return &v1.ListResourceRequestsResponse{
		ResourceRequestStatusSummaries: progressEvents,
	}, nil
}

func buildProgressEvent(req *types.Request) *v1.ProgressEvent {
	return &v1.ProgressEvent{
		ErrorCode:       req.ErrorCode,
		EventTime:       req.RequestTime,
		Identifier:      req.Identifier.String(),
		Operation:       req.Operation,
		OperationStatus: req.OperationStatus,
		RequestToken:    req.Id,
		ResourceModel:   req.DesiredState,
		RetryAfter:      req.RetryAfter,
		StatusMessage:   req.StatusMessage,
		TypeName:        req.Identifier.ResourceType,
	}
}

func buildResourceDescription(rs *types.Resource) *v1.ResourceDescription {
	return &v1.ResourceDescription{
		Identifier: rs.Identifier.String(),
		Properties: rs.State,
		CreatedAt:  rs.CreatedAt.Format(time.RFC3339),
	}
}
