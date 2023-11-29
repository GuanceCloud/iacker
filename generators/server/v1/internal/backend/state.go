package backend

import (
	"context"
	"fmt"

	"github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1/errors"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/provider"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
	krerrors "github.com/go-kratos/kratos/v2/errors"
)

type stateFunc func(ctx context.Context, d *types.Request) (stateFunc, error)

func (w *StateBackend) apply(ctx context.Context, d *types.Request) (*types.Request, error) {
	var err error
	var nextState stateFunc = w.statePending
	for {
		nextState, err = nextState(ctx, d)
		if nextState == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("apply failed: %w", err)
	}
	return w.GetRequestStatus(ctx, d.Id)
}

func (w *StateBackend) statePending(ctx context.Context, d *types.Request) (stateFunc, error) {
	// TODO: hash and schedule the task
	// Default is scheduled to current node now
	return w.stateInProgress, nil
}

func (w *StateBackend) stateInProgress(ctx context.Context, d *types.Request) (stateFunc, error) {
	var err error
	rs, ok := w.Resources[d.Identifier.ResourceType]
	if !ok {
		return w.stateFail(errors.ErrorServiceInternalError("invalid resource type: %v", d.Identifier.ResourceType))
	}

	// TODO: lock resource and use transaction propagation
	d.OperationStatus = types.RequestStatusInProgress
	if err := w.Request.Save(ctx, d); err != nil {
		return nil, errors.ErrorServiceInternalError("save request status to %q failed: %s", d.OperationStatus, err)
	}

	desiredState := &provider.State{
		Config:     rs.GetConfig(),
		Identifier: d.Identifier,
		State:      d.DesiredResource().State,
	}
	switch d.Operation {
	case types.RequestOperationCreate:
		err = rs.Create(ctx, desiredState)
	case types.RequestOperationUpdate:
		err = rs.Update(ctx, desiredState)
	case types.RequestOperationDelete:
		err = rs.Delete(ctx, desiredState)
	default:
		return nil, errors.ErrorServiceInternalError("unknown operation %q", d.Operation)
	}
	if err != nil {
		return w.stateFail(errors.ErrorServiceInternalError("reconcile resource %q failed: %s", d.Identifier, err))
	}
	d.Identifier = desiredState.Identifier
	return func(ctx context.Context, _ *types.Request) (stateFunc, error) {
		return w.stateSuccess(ctx, d)
	}, nil
}

func (w *StateBackend) stateSuccess(ctx context.Context, d *types.Request) (stateFunc, error) {
	if err := w.Resource.Save(ctx, d.DesiredResource()); err != nil {
		return nil, errors.ErrorServiceInternalError("save resource %q failed: %s", d.Identifier, err)
	}

	d.OperationStatus = types.RequestStatusSuccess
	if err := w.Request.Save(ctx, d); err != nil {
		return nil, errors.ErrorServiceInternalError("save request status to %q failed: %s", d.OperationStatus, err)
	}
	// TODO: unlock resource
	return nil, nil
}

//nolint:unused
func (w *StateBackend) stateCancelInProgress(ctx context.Context, d *types.Request) (stateFunc, error) {
	d.OperationStatus = types.RequestStatusCancelInProgress
	if err := w.Request.Save(ctx, d); err != nil {
		return nil, errors.ErrorServiceInternalError("save request status to %q failed: %s", d.OperationStatus, err)
	}
	return w.stateCancelComplete, nil
}

//nolint:unused
func (w *StateBackend) stateCancelComplete(ctx context.Context, d *types.Request) (stateFunc, error) { //
	d.OperationStatus = types.RequestStatusCancelComplete
	if err := w.Request.Save(ctx, d); err != nil {
		return nil, errors.ErrorServiceInternalError("save request status to %q failed: %s", d.OperationStatus, err)
	}
	// TODO: unlock resource
	return nil, nil
}

func (w *StateBackend) stateFail(err *krerrors.Error) (stateFunc, error) {
	return func(ctx context.Context, d *types.Request) (stateFunc, error) {
		d.StatusMessage = err.Error()
		d.OperationStatus = types.RequestStatusFailed
		d.ErrorCode = fmt.Sprint(err.Code)
		if dbErr := w.Request.Save(ctx, d); dbErr != nil {
			return nil, errors.ErrorServiceInternalError("cannot save the failed status: %s; db: %s", err, dbErr)
		}
		// TODO: unlock resource
		return nil, nil
	}, nil
}
