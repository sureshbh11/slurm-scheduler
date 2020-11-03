//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package abregistrar is a generated GoMock package.
package abregistrar

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/hpe-hcss/iam-lib/pkg/ab-registration-client"
	resource "github.com/hpe-hcss/iam-lib/pkg/resource"
	reflect "reflect"
)

// MockRegistrar is a mock of Registrar interface
type MockRegistrar struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrarMockRecorder
}

// MockRegistrarMockRecorder is the mock recorder for MockRegistrar
type MockRegistrarMockRecorder struct {
	mock *MockRegistrar
}

// NewMockRegistrar creates a new mock instance
func NewMockRegistrar(ctrl *gomock.Controller) *MockRegistrar {
	mock := &MockRegistrar{ctrl: ctrl}
	mock.recorder = &MockRegistrarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrar) EXPECT() *MockRegistrarMockRecorder {
	return m.recorder
}

// RegisterAuthorizationBroker mocks base method
func (m *MockRegistrar) RegisterAuthorizationBroker(ctx context.Context, tenantID, token, spaceID, externalURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAuthorizationBroker", ctx, tenantID, token, spaceID, externalURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAuthorizationBroker indicates an expected call of RegisterAuthorizationBroker
func (mr *MockRegistrarMockRecorder) RegisterAuthorizationBroker(ctx, tenantID, token, spaceID, externalURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAuthorizationBroker", reflect.TypeOf((*MockRegistrar)(nil).RegisterAuthorizationBroker), ctx, tenantID, token, spaceID, externalURL)
}

// AddRoles mocks base method
func (m *MockRegistrar) AddRoles(ctx context.Context, tenantID, token, spaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoles", ctx, tenantID, token, spaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoles indicates an expected call of AddRoles
func (mr *MockRegistrarMockRecorder) AddRoles(ctx, tenantID, token, spaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoles", reflect.TypeOf((*MockRegistrar)(nil).AddRoles), ctx, tenantID, token, spaceID)
}

// DeRegisterAuthorizationBroker mocks base method
func (m *MockRegistrar) DeRegisterAuthorizationBroker(ctx context.Context, tenantID, token, spaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeRegisterAuthorizationBroker", ctx, tenantID, token, spaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeRegisterAuthorizationBroker indicates an expected call of DeRegisterAuthorizationBroker
func (mr *MockRegistrarMockRecorder) DeRegisterAuthorizationBroker(ctx, tenantID, token, spaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeRegisterAuthorizationBroker", reflect.TypeOf((*MockRegistrar)(nil).DeRegisterAuthorizationBroker), ctx, tenantID, token, spaceID)
}

// DeleteRoles mocks base method
func (m *MockRegistrar) DeleteRoles(ctx context.Context, tenantID, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoles", ctx, tenantID, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoles indicates an expected call of DeleteRoles
func (mr *MockRegistrarMockRecorder) DeleteRoles(ctx, tenantID, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoles", reflect.TypeOf((*MockRegistrar)(nil).DeleteRoles), ctx, tenantID, token)
}

// MockRegistrarClient is a mock of RegistrarClient interface
type MockRegistrarClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrarClientMockRecorder
}

// MockRegistrarClientMockRecorder is the mock recorder for MockRegistrarClient
type MockRegistrarClientMockRecorder struct {
	mock *MockRegistrarClient
}

// NewMockRegistrarClient creates a new mock instance
func NewMockRegistrarClient(ctrl *gomock.Controller) *MockRegistrarClient {
	mock := &MockRegistrarClient{ctrl: ctrl}
	mock.recorder = &MockRegistrarClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrarClient) EXPECT() *MockRegistrarClientMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockRegistrarClient) Register(arg0 context.Context, arg1 client.ClientDetails, arg2 resource.AuthorizationBroker, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockRegistrarClientMockRecorder) Register(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistrarClient)(nil).Register), arg0, arg1, arg2, arg3)
}

// AddRole mocks base method
func (m *MockRegistrarClient) AddRole(arg0 context.Context, arg1 client.ClientDetails, arg2 string, arg3 resource.Role, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRole", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRole indicates an expected call of AddRole
func (mr *MockRegistrarClientMockRecorder) AddRole(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRole", reflect.TypeOf((*MockRegistrarClient)(nil).AddRole), arg0, arg1, arg2, arg3, arg4)
}

// DeleteRole mocks base method
func (m *MockRegistrarClient) DeleteRole(arg0 context.Context, arg1 client.ClientDetails, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole
func (mr *MockRegistrarClientMockRecorder) DeleteRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRegistrarClient)(nil).DeleteRole), arg0, arg1, arg2)
}
