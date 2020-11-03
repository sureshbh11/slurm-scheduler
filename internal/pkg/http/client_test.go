// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/iam"
)

type mockRequest struct {
	Name string
}

func Test_NewHTTPRequest_(t *testing.T) {
	httpClient := NewClient()
	assert.NotNil(t, httpClient)

	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockURL := "http://127.0.0.1:8080"

	requestBytes, _ := json.Marshal(mockRequest{Name: "test-http-request"})
	mockRequest := bytes.NewReader(requestBytes)
	response, err := NewHTTPRequest(mockContext, "GET", mockURL, mockRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	response, err = NewHTTPRequest(nil, "GET", mockURL, mockRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	mockHttpclient := common.NewMockHTTPClient(ctrl)
	mockHTTPResp, _ := iam.NewJSONResponse(200, "")
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	resResponse, err := ExecuteHTTPRequest(mockContext, mockHttpclient, mockURL, "GET",
		headers, mockRequest)
	assert.Equal(t, "200", resResponse.Status)
	assert.Nil(t, err)

	queryMap := map[string]string{}
	queryMap["filter"] = "mockQuery"
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	resResponse, err = ExecuteHTTPRequestWithQuery(mockContext, mockHttpclient, mockURL,
		queryMap, "GET", headers, mockRequest)
	assert.Nil(t, err)

	errMsg := "HTTP Failure"
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, errors.New(errMsg)).Times(1)
	resResponse, err = ExecuteHTTPRequestWithQuery(mockContext, mockHttpclient, mockURL,
		queryMap, "GET", headers, mockRequest)
	assert.NotNil(t, err)
	assert.Nil(t, resResponse)
	assert.Contains(t, err.Error(), errMsg)

	mockHTTPResp.StatusCode = 302
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	resResponse, err = ExecuteHTTPRequestWithQuery(mockContext, mockHttpclient, mockURL,
		queryMap, "GET", headers, mockRequest)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "302")

	mockHttpRequest, err := http.NewRequest("GET", "http://google.com", mockRequest)
	loggerObj := NewLoggerTransport(http.DefaultTransport)
	loggerObj.RoundTrip(mockHttpRequest)

}
