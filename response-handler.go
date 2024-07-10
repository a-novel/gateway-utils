package gatewayutils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrNoResponseBody = errors.New("empty response")
)

// GetResponseError extract the body of an HTTP response with a fail status code, so it can be properly displayed in
// the error message.
//
// If there is no response body, ErrNoResponseBody is thrown.
func GetResponseError(res *http.Response) error {
	if res.Body == nil {
		return ErrNoResponseBody
	}

	strResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return errors.New(string(strResponse))
}

// ExtractJSONResponse decodes the body of a HTTP response into a given struct.
func ExtractJSONResponse(res *http.Response, dest interface{}) error {
	if err := json.NewDecoder(res.Body).Decode(dest); err != nil {
		return err
	}

	return nil
}
