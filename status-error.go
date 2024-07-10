package gatewayutils

import (
	"fmt"
	"net/http"
)

// StatusError is thrown when an API call returns a non-success status.
type StatusError struct {
	// Status is the status returned by the failed API call.
	Status int
}

// Error returns the content of a StatusError as a string. Complies with the Error interface.
func (e *StatusError) Error() string {
	return fmt.Sprintf("unexpected status: %v", e.Status)
}

// EnsureStatus throws an error if the response status is not among a list of expected statuses.
func EnsureStatus(res *http.Response, expected ...int) error {
	for _, status := range expected {
		// We found the status in the list of expected statuses. The response is valid.
		if res.StatusCode == status {
			return nil
		}
	}

	// There was no match in the list of expected statuses. The response is invalid.
	return &StatusError{Status: res.StatusCode}
}
