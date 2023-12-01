package errors

import (
	fmt "fmt"

	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

const (
	// ErrorReasonAccessDenied is AccessDenied is returned when the request is not authorized.
	ErrorReasonAccessDenied = "AccessDenied"

	// ErrorReasonAlreadyExists is AlreadyExists is returned when the resource already exists.
	ErrorReasonAlreadyExists = "AlreadyExists"

	// ErrorReasonGeneralServiceException is NotStabilized is returned when the error is uncategorized.
	ErrorReasonGeneralServiceException = "GeneralServiceException"

	// ErrorReasonInvalidCredentials is InvalidCredentials is returned when the request is not authorized.
	ErrorReasonInvalidCredentials = "InvalidCredentials"

	// ErrorReasonInvalidRequest is InvalidRequest is returned when the request is invalid.
	ErrorReasonInvalidRequest = "InvalidRequest"

	// ErrorReasonNetworkFailure is NetworkFailure is returned when the network is not available.
	ErrorReasonNetworkFailure = "NetworkFailure"

	// ErrorReasonNotFound is NotFound is returned when the resource is not found.
	ErrorReasonNotFound = "NotFound"

	// ErrorReasonNotUpdatable is NotUpdatable is returned when the request is trying to update a field that is not updatable.
	ErrorReasonNotUpdatable = "NotUpdatable"

	// ErrorReasonResourceConflict is ResourceConflict is returned when the resource is in conflict.
	ErrorReasonResourceConflict = "ResourceConflict"

	// ErrorReasonServiceInternalError is ServiceInternalError is returned when the error is unknown.
	ErrorReasonServiceInternalError = "ServiceInternalError"

	// ErrorReasonServiceLimitExceeded is ServiceLimitExceeded is returned when the request is reach the service limit.
	ErrorReasonServiceLimitExceeded = "ServiceLimitExceeded"

	// ErrorReasonServiceTimeout is ServiceTimeout is returned when the request is timeout.
	ErrorReasonServiceTimeout = "ServiceTimeout"

	// ErrorReasonThrottling is Throttling is returned when the request is throttled.
	ErrorReasonThrottling = "Throttling"

	// ErrorReasonUnspecified is Unspecified is returned when the error is unknown.
	ErrorReasonUnspecified = "Unspecified"
)

// IsAccessDenied is returned when AccessDenied is returned when the request is not authorized.
func IsAccessDenied(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonAccessDenied && e.Code == 403
}

// ErrorAccessDenied is returned when AccessDenied is returned when the request is not authorized.
func ErrorAccessDenied(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReasonAccessDenied, fmt.Sprintf(format, args...))
}

// IsAlreadyExists is returned when AlreadyExists is returned when the resource already exists.
func IsAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonAlreadyExists && e.Code == 409
}

// ErrorAlreadyExists is returned when AlreadyExists is returned when the resource already exists.
func ErrorAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReasonAlreadyExists, fmt.Sprintf(format, args...))
}

// IsGeneralServiceException is returned when NotStabilized is returned when the error is uncategorized.
func IsGeneralServiceException(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonGeneralServiceException && e.Code == 500
}

// ErrorGeneralServiceException is returned when NotStabilized is returned when the error is uncategorized.
func ErrorGeneralServiceException(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReasonGeneralServiceException, fmt.Sprintf(format, args...))
}

// IsInvalidCredentials is returned when InvalidCredentials is returned when the request is not authorized.
func IsInvalidCredentials(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonInvalidCredentials && e.Code == 403
}

// ErrorInvalidCredentials is returned when InvalidCredentials is returned when the request is not authorized.
func ErrorInvalidCredentials(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReasonInvalidCredentials, fmt.Sprintf(format, args...))
}

// IsInvalidRequest is returned when InvalidRequest is returned when the request is invalid.
func IsInvalidRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonInvalidRequest && e.Code == 400
}

// ErrorInvalidRequest is returned when InvalidRequest is returned when the request is invalid.
func ErrorInvalidRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReasonInvalidRequest, fmt.Sprintf(format, args...))
}

// IsNetworkFailure is returned when NetworkFailure is returned when the network is not available.
func IsNetworkFailure(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonNetworkFailure && e.Code == 503
}

// ErrorNetworkFailure is returned when NetworkFailure is returned when the network is not available.
func ErrorNetworkFailure(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ErrorReasonNetworkFailure, fmt.Sprintf(format, args...))
}

// IsNotFound is returned when NotFound is returned when the resource is not found.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonNotFound && e.Code == 404
}

// ErrorNotFound is returned when NotFound is returned when the resource is not found.
func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReasonNotFound, fmt.Sprintf(format, args...))
}

// IsNotUpdatable is returned when NotUpdatable is returned when the request is trying to update a field that is not updatable.
func IsNotUpdatable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonNotUpdatable && e.Code == 400
}

// ErrorNotUpdatable is returned when NotUpdatable is returned when the request is trying to update a field that is not updatable.
func ErrorNotUpdatable(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReasonNotUpdatable, fmt.Sprintf(format, args...))
}

// IsResourceConflict is returned when ResourceConflict is returned when the resource is in conflict.
func IsResourceConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonResourceConflict && e.Code == 409
}

// ErrorResourceConflict is returned when ResourceConflict is returned when the resource is in conflict.
func ErrorResourceConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReasonResourceConflict, fmt.Sprintf(format, args...))
}

// IsServiceInternalError is returned when ServiceInternalError is returned when the error is unknown.
func IsServiceInternalError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonServiceInternalError && e.Code == 500
}

// ErrorServiceInternalError is returned when ServiceInternalError is returned when the error is unknown.
func ErrorServiceInternalError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReasonServiceInternalError, fmt.Sprintf(format, args...))
}

// IsServiceLimitExceeded is returned when ServiceLimitExceeded is returned when the request is reach the service limit.
func IsServiceLimitExceeded(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonServiceLimitExceeded && e.Code == 429
}

// ErrorServiceLimitExceeded is returned when ServiceLimitExceeded is returned when the request is reach the service limit.
func ErrorServiceLimitExceeded(format string, args ...interface{}) *errors.Error {
	return errors.New(429, ErrorReasonServiceLimitExceeded, fmt.Sprintf(format, args...))
}

// IsServiceTimeout is returned when ServiceTimeout is returned when the request is timeout.
func IsServiceTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonServiceTimeout && e.Code == 504
}

// ErrorServiceTimeout is returned when ServiceTimeout is returned when the request is timeout.
func ErrorServiceTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, ErrorReasonServiceTimeout, fmt.Sprintf(format, args...))
}

// IsThrottling is returned when Throttling is returned when the request is throttled.
func IsThrottling(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonThrottling && e.Code == 429
}

// ErrorThrottling is returned when Throttling is returned when the request is throttled.
func ErrorThrottling(format string, args ...interface{}) *errors.Error {
	return errors.New(429, ErrorReasonThrottling, fmt.Sprintf(format, args...))
}

// IsUnspecified is returned when Unspecified is returned when the error is unknown.
func IsUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReasonUnspecified && e.Code == 500
}

// ErrorUnspecified is returned when Unspecified is returned when the error is unknown.
func ErrorUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReasonUnspecified, fmt.Sprintf(format, args...))
}
