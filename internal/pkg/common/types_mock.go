// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./types.go

package common

import (
	http "net/http"

	gomock "github.com/golang/mock/gomock"
)

// Mock of HTTPClient interface
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *_MockHTTPClientRecorder
}

// Recorder for MockHTTPClient (not exported)
type _MockHTTPClientRecorder struct {
	mock *MockHTTPClient
}

func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &_MockHTTPClientRecorder{mock}
	return mock
}

func (_m *MockHTTPClient) EXPECT() *_MockHTTPClientRecorder {
	return _m.recorder
}

func (_m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHTTPClientRecorder) Do(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Do", arg0)
}