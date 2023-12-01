package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/avast/retry-go"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/hashicorp/go-multierror"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	v1 "github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

type Client interface {
	CreateResource(ctx context.Context, req *v1.CreateResourceRequest, opts ...Option) (rsp *v1.CreateResourceResponse, err error)
	DeleteResource(ctx context.Context, req *v1.DeleteResourceRequest, opts ...Option) (rsp *v1.DeleteResourceResponse, err error)
	GetResource(ctx context.Context, req *v1.GetResourceRequest, opts ...Option) (rsp *v1.GetResourceResponse, err error)
	ListResources(ctx context.Context, req *v1.ListResourcesRequest, opts ...Option) (rsp *v1.ListResourcesResponse, err error)
	UpdateResource(ctx context.Context, req *v1.UpdateResourceRequest, opts ...Option) (rsp *v1.UpdateResourceResponse, err error)
}

type Option func(Client) error

func NewClient(options ...Option) (Client, error) {
	c := &client{}
	var mErr error
	for _, option := range options {
		err := option(c)
		if err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	return c, mErr
}

type client struct {
	token                string
	region               string
	endpoint             string
	wait                 bool
	maxRetries           int
	nonIdempotentRetried bool
}

// WithEndpoint is an option to set the endpoint of the service.
func WithEndpoint(endpoint string) Option {
	return func(c Client) error {
		c.(*client).endpoint = endpoint
		return nil
	}
}

// WithWait is an option to set if need to wait the request is completed.
func WithWait(wait bool) Option {
	return func(c Client) error {
		c.(*client).wait = wait
		return nil
	}
}

// WithRegion is an option to set the region of the service.
func WithRegion(region string) Option {
	return func(c Client) error {
		c.(*client).region = region
		// c.(*client).endpoint = "https://cloudcontrol.guance.io"
		c.(*client).endpoint = "http://localhost:8000"
		return nil
	}
}

// WithAccessToken is an option to set the access token of the service.
func WithAccessToken(token string) Option {
	return func(c Client) error {
		c.(*client).token = token
		return nil
	}
}

// WithMaxRetries is an option to set the max retries of the request.
func WithMaxRetries(maxRetries int) Option {
	return func(c Client) error {
		c.(*client).maxRetries = maxRetries
		return nil
	}
}

// WithNonIdempotentRetried is an option to set if need to retry the non idempotent request.
func WithNonIdempotentRetried(nonIdempotentRetried bool) Option {
	return func(c Client) error {
		c.(*client).nonIdempotentRetried = nonIdempotentRetried
		return nil
	}
}

func (c client) clone() client {
	return client{
		token:                c.token,
		region:               c.region,
		endpoint:             c.endpoint,
		wait:                 c.wait,
		maxRetries:           c.maxRetries,
		nonIdempotentRetried: c.nonIdempotentRetried,
	}
}

func (c client) withOptions(opts ...Option) (*client, error) {
	var mErr error
	cc := c.clone()
	for _, opt := range opts {
		if err := opt(&cc); err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	return &cc, mErr
}

func (c client) CreateResource(ctx context.Context, req *v1.CreateResourceRequest, opts ...Option) (rsp *v1.CreateResourceResponse, err error) {
	if !c.nonIdempotentRetried {
		opts = append(opts, WithMaxRetries(0))
	}

	cc, err := c.withOptions(opts...)
	if err != nil {
		return nil, err
	}

	req.Region = cc.region
	req.AccessToken = cc.token

	// request to create resource
	rsp = &v1.CreateResourceResponse{}
	err = cc.invoke(ctx, "/cloud-control/create-resource", req, rsp)
	if err != nil {
		return nil, err
	}

	if !cc.wait {
		return rsp, nil
	}

	if err := c.Wait(ctx, WaitOptions{RequestId: rsp.ProgressEvent.RequestToken}); err != nil {
		return nil, fmt.Errorf("failed when retring to create resource: %w", err)
	}
	return rsp, nil
}

func (c client) DeleteResource(ctx context.Context, req *v1.DeleteResourceRequest, opts ...Option) (rsp *v1.DeleteResourceResponse, err error) {
	if !c.nonIdempotentRetried {
		opts = append(opts, WithMaxRetries(0))
	}

	cc, err := c.withOptions(opts...)
	if err != nil {
		return nil, err
	}

	req.AccessToken = cc.token

	// request to create resource
	rsp = &v1.DeleteResourceResponse{}
	err = cc.invoke(ctx, "/cloud-control/delete-resource", req, rsp)
	if err != nil {
		return nil, err
	}

	if !cc.wait {
		return rsp, nil
	}

	if err := c.Wait(ctx, WaitOptions{RequestId: rsp.ProgressEvent.RequestToken}); err != nil {
		return nil, fmt.Errorf("failed when retring to create resource: %w", err)
	}
	return rsp, nil
}

func (c client) GetResource(ctx context.Context, req *v1.GetResourceRequest, opts ...Option) (rsp *v1.GetResourceResponse, err error) {
	cc, err := c.withOptions(opts...)
	if err != nil {
		return nil, err
	}

	req.AccessToken = cc.token

	// request to create resource
	rsp = &v1.GetResourceResponse{}
	err = cc.invoke(ctx, "/cloud-control/get-resource", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c client) ListResources(ctx context.Context, req *v1.ListResourcesRequest, opts ...Option) (rsp *v1.ListResourcesResponse, err error) {
	cc, err := c.withOptions(opts...)
	if err != nil {
		return nil, err
	}

	req.AccessToken = cc.token

	// request to create resource
	rsp = &v1.ListResourcesResponse{}
	err = cc.invoke(ctx, "/cloud-control/list-resources", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c client) UpdateResource(ctx context.Context, req *v1.UpdateResourceRequest, opts ...Option) (rsp *v1.UpdateResourceResponse, err error) {
	if !c.nonIdempotentRetried {
		opts = append(opts, WithMaxRetries(0))
	}

	cc, err := c.withOptions(opts...)
	if err != nil {
		return nil, err
	}

	req.AccessToken = cc.token

	// request to create resource
	rsp = &v1.UpdateResourceResponse{}
	err = cc.invoke(ctx, "/cloud-control/update-resource", req, rsp)
	if err != nil {
		return nil, err
	}

	if !cc.wait {
		return rsp, nil
	}

	if err := c.Wait(ctx, WaitOptions{RequestId: rsp.ProgressEvent.RequestToken}); err != nil {
		return nil, fmt.Errorf("failed when retring to create resource: %w", err)
	}
	return rsp, nil
}

func (c client) invoke(ctx context.Context, path string, in interface{}, out proto.Message) error {
	reqBytes, err := json.Marshal(in)
	if err != nil {
		return err
	}

	httpClient := &http.Client{}

	// Send request
	resp, err := httpClient.Post(
		fmt.Sprintf("%s/%s", strings.Trim(c.endpoint, "/"), strings.Trim(path, "/")),
		"application/json",
		bytes.NewBuffer(reqBytes),
	)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return err
	}

	// Got errors
	if resp.StatusCode >= 400 {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read response body failed with status code %d", resp.StatusCode)
		}
		fmt.Println("[Got]", path, string(respBytes))
		outErr := &errors.Error{}
		if err := protojson.Unmarshal(respBytes, outErr); err != nil {
			return fmt.Errorf("unmarshal failed with status code %d: %s", resp.StatusCode, string(respBytes))
		}
		return outErr
	}

	// Got response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("[Got]", path, string(respBytes))
	if err := protojson.Unmarshal(respBytes, out); err != nil {
		return err
	}
	return nil
}

type WaitOptions struct {
	RequestId string
}

func (c client) Wait(ctx context.Context, opts WaitOptions) error {
	return retry.Do(
		func() error {
			// wait for the request is completed
			r, err := c.getResourceRequestStatus(ctx, &v1.GetResourceRequestStatusRequest{
				RequestToken: opts.RequestId,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Waiting for request to be completed, current status: %s\n", r.ProgressEvent.OperationStatus)

			switch r.ProgressEvent.OperationStatus {
			case types.RequestStatusPending,
				types.RequestStatusInProgress,
				types.RequestStatusCancelInProgress:
				return fmt.Errorf("request is not completed")
			case types.RequestStatusSuccess,
				types.RequestStatusCancelComplete:
				return nil
			case types.RequestStatusFailed:
				return fmt.Errorf("failed: %s", r.ProgressEvent.StatusMessage)
			default:
				return fmt.Errorf("unknown status %s", r.ProgressEvent.OperationStatus)
			}
		},
		retry.RetryIf(func(err error) bool {
			return strings.Contains(err.Error(), "request is not completed")
		}),
	)
}

func (c client) getResourceRequestStatus(ctx context.Context, g *v1.GetResourceRequestStatusRequest) (*v1.GetResourceRequestStatusResponse, error) {
	rsp := &v1.GetResourceRequestStatusResponse{}
	err := c.invoke(ctx, "/cloud-control/get-resource-request-status", g, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
