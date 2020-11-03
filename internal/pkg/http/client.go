// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package http

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hpe-hcss/loglib/pkg/errors"

	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
)

//NewClient returns http client
func NewClient() *http.Client {
	return &http.Client{
		Transport: NewLoggerTransport(http.DefaultTransport),
		Timeout:   10 * time.Second,
	}
}

//NewHTTPRequest returns http request object with given context, if context is not nil
func NewHTTPRequest(
	ctx context.Context,
	method, url string,
	body io.Reader,
) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		return req, nil
	}
	return req.WithContext(ctx), nil
}

// ExecuteHTTPRequest execute request using given client
func ExecuteHTTPRequest(
	ctx context.Context,
	client common.HTTPClient,
	url string,
	method string,
	headers map[string]string,
	body io.Reader,
) (*http.Response, error) {
	return ExecuteHTTPRequestWithQuery(ctx, client, url, nil, method, headers, body)
}

// ExecuteHTTPRequestWithQuery execute request using given client with query map
func ExecuteHTTPRequestWithQuery(
	ctx context.Context,
	client common.HTTPClient,
	url string,
	queryMap map[string]string,
	method string,
	headers map[string]string,
	body io.Reader,
) (*http.Response, error) {
	req, err := NewHTTPRequest(ctx, method, url, body)
	if err != nil {
		return nil, errors.Wrapf(err, "creating request for : [%s]", url)
	}

	q := req.URL.Query()
	for key, value := range queryMap {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "getting response for: [%s]", url)
	}
	defer res.Body.Close()

	responseBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, errors.Wrapf(readErr, "reading response body for: [%s]", url)
	}

	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))

	if res.StatusCode > 299 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "while reading response body, StatusCode: %d", res.StatusCode)
		}
		return nil, &HTTPError{
			StatusCode: res.StatusCode,
			Body:       string(body),
		}
	}

	return res, nil
}
