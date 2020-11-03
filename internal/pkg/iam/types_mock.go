//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	authorization "github.com/hpe-hcss/iam-lib/pkg/resource/authorization"
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

// MockAuthorizer is a mock of Authorizer interface
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// Authorize mocks base method
func (m *MockAuthorizer) Authorize(ctx context.Context, tenant, permission, userID string, resource authorization.ResourceIdentifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, tenant, permission, userID, resource)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorize indicates an expected call of Authorize
func (mr *MockAuthorizerMockRecorder) Authorize(ctx, tenant, permission, userID, resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthorizer)(nil).Authorize), ctx, tenant, permission, userID, resource)
}
