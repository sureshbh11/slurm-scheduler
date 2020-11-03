//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package http

import (
	"net/http"

	"github.com/hpe-hcss/loglib/pkg/log"
)

//LoggerTransport wraps und tripper interface for logging
type LoggerTransport struct {
	tr http.RoundTripper
}

//NewLoggerTransport wraps roundtripper interface around LoggerTransport
func NewLoggerTransport(tr http.RoundTripper) http.RoundTripper {
	return &LoggerTransport{tr: tr}
}

//RoundTrip logs method and path before calling nested round tripper interface
func (trans *LoggerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Infof(req.Context(), "method: [%s], path:[%s]", req.Method, req.URL.String())
	resp, err := trans.tr.RoundTrip(req)
	if resp != nil {
		if resp.StatusCode < 399 {
			log.Infof(req.Context(),
				"method: [%s], path:[%s], Response status: [%d],",
				req.Method, req.URL.String(), resp.StatusCode,
			)
		} else {
			log.Errorf(req.Context(),
				"method: [%s], path:[%s], Response status: [%d],",
				req.Method, req.URL.String(), resp.StatusCode,
			)
		}
	}
	return resp, err
}
