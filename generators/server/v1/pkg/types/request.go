package types

import (
	"time"

	"github.com/google/uuid"
)

const (
	// RequestOperationCreate is the operation type for create.
	RequestOperationCreate = "CREATE"

	// RequestOperationDelete is the operation type for delete.
	RequestOperationDelete = "DELETE"

	// RequestOperationUpdate is the operation type for update.
	RequestOperationUpdate = "UPDATE"
)

const (
	// RequestStatusPending is the status for pending.
	RequestStatusPending = "PENDING"

	// RequestStatusInProgress is the status for in progress.
	RequestStatusInProgress = "IN_PROGRESS"

	// RequestStatusSuccess is the status for success.
	RequestStatusSuccess = "SUCCESS"

	// RequestStatusFailed is the status for failed.
	RequestStatusFailed = "FAILED"

	// RequestStatusCancelInProgress is the status for cancel in progress.
	RequestStatusCancelInProgress = "CANCEL_IN_PROGRESS"

	// RequestStatusCancelComplete is the status for cancel complete.
	RequestStatusCancelComplete = "CANCEL_COMPLETE"
)

// Request is the current status of a resource operation request.
type Request struct {
	// For requests with a status of FAILED, the associated error code.
	//
	// Valid Values:
	// NotUpdatable
	// | InvalidRequest
	// | AccessDenied
	// | InvalidCredentials
	// | AlreadyExists
	// | NotFound
	// | ResourceConflict
	// | Throttling
	// | ServiceLimitExceeded
	// | NotStabilized
	// | GeneralServiceException
	// | ServiceInternalError
	// | ServiceTimeout
	// | NetworkFailure
	// | InternalFailure
	ErrorCode string

	// When the resource operation request was initiated.
	RequestTime int64

	// The primary identifier for the resource.
	Identifier *Identifier

	// The resource operation type.
	//
	// Valid Values:
	// CREATE
	// | DELETE
	// | UPDATE
	Operation string

	// The current status of the resource operation request.
	//
	// PENDING: The resource operation hasn't yet started.
	// IN_PROGRESS: The resource operation is currently in progress.
	// SUCCESS: The resource operation has successfully completed.
	// FAILED: The resource operation has failed. Refer to the error code and status message for more information.
	// CANCEL_IN_PROGRESS: The resource operation is in the process of being canceled.
	// CANCEL_COMPLETE: The resource operation has been canceled.
	OperationStatus string

	// The unique token representing this resource operation request.
	// Use the RequestToken with GetResourceRequestStatus to return the current status of a resource operation request.
	Id string

	// A JSON string containing the resource model,
	// consisting of each resource property and its current value.
	DesiredState string

	// When to next request the status of this resource operation request.
	RetryAfter int64

	// Any message explaining the current status.
	StatusMessage string
}

func (r *Request) DesiredResource() *Resource {
	return &Resource{
		Identifier: r.Identifier,
		State:      r.DesiredState,
	}
}

func NewRequest(rs *Resource, op string) *Request {
	return &Request{
		Id:              uuid.New().String(),
		Identifier:      rs.Identifier,
		DesiredState:    rs.State,
		Operation:       op,
		OperationStatus: RequestStatusPending,
		RequestTime:     time.Now().UnixNano(),
	}
}
