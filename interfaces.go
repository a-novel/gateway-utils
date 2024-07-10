package gatewayutils

import (
	"context"
	"errors"
)

var (
	// ErrUnavailable is thrown when a ping to am API does not return anything - meaning the service is effectively
	// down in the current environment.
	ErrUnavailable = errors.New("failed to reach the api")
)

// PingAPI represents a standard interface for an API healthcheck.
//
// If the API is unavailable, ErrUnavailable should be returned. Any other error indicates a critical situation.
//
// A ping API should always forward the response status.
type PingAPI interface {
	Call(ctx context.Context) (int, error)
}
