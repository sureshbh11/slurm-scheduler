// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hpe-hcss/scm-lib/pkg/scmclient (interfaces: SCMAPI)

// Package authzbroker is a generated GoMock package.
package authzbroker

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	identityclient "github.com/hpe-hcss/iam-lib/pkg/identity-client"
	scmclient "github.com/hpe-hcss/scm-lib/pkg/scmclient"
	reflect "reflect"
)

// MockSCMAPI is a mock of SCMAPI interface
type MockSCMAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSCMAPIMockRecorder
}

// MockSCMAPIMockRecorder is the mock recorder for MockSCMAPI
type MockSCMAPIMockRecorder struct {
	mock *MockSCMAPI
}

// NewMockSCMAPI creates a new mock instance
func NewMockSCMAPI(ctrl *gomock.Controller) *MockSCMAPI {
	mock := &MockSCMAPI{ctrl: ctrl}
	mock.recorder = &MockSCMAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSCMAPI) EXPECT() *MockSCMAPIMockRecorder {
	return m.recorder
}

// CaptureTenantClient mocks base method
func (m *MockSCMAPI) CaptureTenantClient(arg0 context.Context, arg1, arg2 string, arg3 scmclient.TenantClient, arg4 string, arg5 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureTenantClient", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureTenantClient indicates an expected call of CaptureTenantClient
func (mr *MockSCMAPIMockRecorder) CaptureTenantClient(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).CaptureTenantClient), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetTenantClient mocks base method
func (m *MockSCMAPI) GetTenantClient(arg0 context.Context, arg1 string) (scmclient.NewTenantClientOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantClient", arg0, arg1)
	ret0, _ := ret[0].(scmclient.NewTenantClientOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenantClient indicates an expected call of GetTenantClient
func (mr *MockSCMAPIMockRecorder) GetTenantClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).GetTenantClient), arg0, arg1)
}

// NewBaseClient mocks base method
func (m *MockSCMAPI) NewBaseClient(arg0 context.Context, arg1, arg2 string) (scmclient.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBaseClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(scmclient.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewBaseClient indicates an expected call of NewBaseClient
func (mr *MockSCMAPIMockRecorder) NewBaseClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBaseClient", reflect.TypeOf((*MockSCMAPI)(nil).NewBaseClient), arg0, arg1, arg2)
}

// RetrieveBaseClients mocks base method
func (m *MockSCMAPI) RetrieveBaseClients(arg0 context.Context) (scmclient.BaseClients, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveBaseClients", arg0)
	ret0, _ := ret[0].(scmclient.BaseClients)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveBaseClients indicates an expected call of RetrieveBaseClients
func (mr *MockSCMAPIMockRecorder) RetrieveBaseClients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveBaseClients", reflect.TypeOf((*MockSCMAPI)(nil).RetrieveBaseClients), arg0)
}

// RetrieveTenantClient mocks base method
func (m *MockSCMAPI) RetrieveTenantClient(arg0 context.Context, arg1 string) (scmclient.TenantClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveTenantClient", arg0, arg1)
	ret0, _ := ret[0].(scmclient.TenantClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveTenantClient indicates an expected call of RetrieveTenantClient
func (mr *MockSCMAPIMockRecorder) RetrieveTenantClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).RetrieveTenantClient), arg0, arg1)
}

// RevokeBaseClient mocks base method
func (m *MockSCMAPI) RevokeBaseClient(arg0 context.Context, arg1 string, arg2 scmclient.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeBaseClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeBaseClient indicates an expected call of RevokeBaseClient
func (mr *MockSCMAPIMockRecorder) RevokeBaseClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeBaseClient", reflect.TypeOf((*MockSCMAPI)(nil).RevokeBaseClient), arg0, arg1, arg2)
}

// RevokeTenantClient mocks base method
func (m *MockSCMAPI) RevokeTenantClient(arg0 context.Context, arg1, arg2 string, arg3 scmclient.TenantClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeTenantClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeTenantClient indicates an expected call of RevokeTenantClient
func (mr *MockSCMAPIMockRecorder) RevokeTenantClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).RevokeTenantClient), arg0, arg1, arg2, arg3)
}

// RotateBaseClient mocks base method
func (m *MockSCMAPI) RotateBaseClient(arg0 context.Context, arg1 string, arg2 scmclient.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateBaseClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RotateBaseClient indicates an expected call of RotateBaseClient
func (mr *MockSCMAPIMockRecorder) RotateBaseClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateBaseClient", reflect.TypeOf((*MockSCMAPI)(nil).RotateBaseClient), arg0, arg1, arg2)
}

// RotateTenantClient mocks base method
func (m *MockSCMAPI) RotateTenantClient(arg0 context.Context, arg1, arg2 string, arg3 scmclient.TenantClient) (scmclient.RotateTenantClientOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateTenantClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(scmclient.RotateTenantClientOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateTenantClient indicates an expected call of RotateTenantClient
func (mr *MockSCMAPIMockRecorder) RotateTenantClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).RotateTenantClient), arg0, arg1, arg2, arg3)
}

// StoreBaseClients mocks base method
func (m *MockSCMAPI) StoreBaseClients(arg0 context.Context, arg1 scmclient.BaseClients) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBaseClients", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreBaseClients indicates an expected call of StoreBaseClients
func (mr *MockSCMAPIMockRecorder) StoreBaseClients(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBaseClients", reflect.TypeOf((*MockSCMAPI)(nil).StoreBaseClients), arg0, arg1)
}

// StoreTenantClient mocks base method
func (m *MockSCMAPI) StoreTenantClient(arg0 context.Context, arg1 string, arg2 scmclient.TenantClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreTenantClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreTenantClient indicates an expected call of StoreTenantClient
func (mr *MockSCMAPIMockRecorder) StoreTenantClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreTenantClient", reflect.TypeOf((*MockSCMAPI)(nil).StoreTenantClient), arg0, arg1, arg2)
}

// VendBaseToken mocks base method
func (m *MockSCMAPI) VendBaseToken(arg0 context.Context, arg1 *identityclient.Client) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VendBaseToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VendBaseToken indicates an expected call of VendBaseToken
func (mr *MockSCMAPIMockRecorder) VendBaseToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VendBaseToken", reflect.TypeOf((*MockSCMAPI)(nil).VendBaseToken), arg0, arg1)
}

// VendTenantToken mocks base method
func (m *MockSCMAPI) VendTenantToken(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VendTenantToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VendTenantToken indicates an expected call of VendTenantToken
func (mr *MockSCMAPIMockRecorder) VendTenantToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VendTenantToken", reflect.TypeOf((*MockSCMAPI)(nil).VendTenantToken), arg0, arg1)
}
