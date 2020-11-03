//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package http

import (
	"fmt"
)

//HTTPError custome http error model
type HTTPError struct {
	Body       string
	StatusCode int
}

//Error implements Error interface
func (e *HTTPError) Error() string {
	return fmt.Sprintf(
		"http request failed with status code: %d, body: %s",
		e.StatusCode,
		e.Body,
	)
}
