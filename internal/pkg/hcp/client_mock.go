// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package hcp is a generated GoMock package.
package hcp

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/hpe-hcss/mlops/internal/pkg/hcp/models"
	models0 "github.com/hpe-hcss/mlops/internal/pkg/models"
	patch "github.com/hpe-hcss/mlops/internal/pkg/patch"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ConfigureAuthInHcp mocks base method
func (m *MockClient) ConfigureAuthInHcp(ctx context.Context, hcpHostURL string, input models0.EpicConfigInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureAuthInHcp", ctx, hcpHostURL, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureAuthInHcp indicates an expected call of ConfigureAuthInHcp
func (mr *MockClientMockRecorder) ConfigureAuthInHcp(ctx, hcpHostURL, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureAuthInHcp", reflect.TypeOf((*MockClient)(nil).ConfigureAuthInHcp), ctx, hcpHostURL, input)
}

// CreateProject mocks base method
func (m *MockClient) CreateProject(ctx context.Context, hcpHostURL string, input models0.HCPCreateProjectInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, hcpHostURL, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockClientMockRecorder) CreateProject(ctx, hcpHostURL, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockClient)(nil).CreateProject), ctx, hcpHostURL, input)
}

// DeleteProject mocks base method
func (m *MockClient) DeleteProject(ctx context.Context, hcpHostURL, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, hcpHostURL, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockClientMockRecorder) DeleteProject(ctx, hcpHostURL, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockClient)(nil).DeleteProject), ctx, hcpHostURL, projectID)
}

// ListProject mocks base method
func (m *MockClient) ListProject(ctx context.Context, hcpHostURL string) ([]models.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProject", ctx, hcpHostURL)
	ret0, _ := ret[0].([]models.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProject indicates an expected call of ListProject
func (mr *MockClientMockRecorder) ListProject(ctx, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProject", reflect.TypeOf((*MockClient)(nil).ListProject), ctx, hcpHostURL)
}

// GetProject mocks base method
func (m *MockClient) GetProject(ctx context.Context, hcpHostURL, projectID string) (models.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", ctx, hcpHostURL, projectID)
	ret0, _ := ret[0].(models.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockClientMockRecorder) GetProject(ctx, hcpHostURL, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockClient)(nil).GetProject), ctx, hcpHostURL, projectID)
}

// UpdateProject mocks base method
func (m *MockClient) UpdateProject(ctx context.Context, hcpHostURL, projectID string, input patch.JSONPatch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, hcpHostURL, projectID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockClientMockRecorder) UpdateProject(ctx, hcpHostURL, projectID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockClient)(nil).UpdateProject), ctx, hcpHostURL, projectID, input)
}

// GetClusters mocks base method
func (m *MockClient) GetClusters(ctx context.Context, hcpHostURL string) (models0.ClusterResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", ctx, hcpHostURL)
	ret0, _ := ret[0].(models0.ClusterResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters
func (mr *MockClientMockRecorder) GetClusters(ctx, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockClient)(nil).GetClusters), ctx, hcpHostURL)
}

// GetModels mocks base method
func (m *MockClient) GetModels(ctx context.Context, hcpHostURL string) (models0.ModelResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModels", ctx, hcpHostURL)
	ret0, _ := ret[0].(models0.ModelResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModels indicates an expected call of GetModels
func (mr *MockClientMockRecorder) GetModels(ctx, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModels", reflect.TypeOf((*MockClient)(nil).GetModels), ctx, hcpHostURL)
}

// GetConfig mocks base method
func (m *MockClient) GetConfig(ctx context.Context, hcpHostURL string) (models.HcpConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, hcpHostURL)
	ret0, _ := ret[0].(models.HcpConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockClientMockRecorder) GetConfig(ctx, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClient)(nil).GetConfig), ctx, hcpHostURL)
}

// ListRoles mocks base method
func (m *MockClient) ListRoles(ctx context.Context, hcpHostURL string) ([]models.RoleResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", ctx, hcpHostURL)
	ret0, _ := ret[0].([]models.RoleResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles
func (mr *MockClientMockRecorder) ListRoles(ctx, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockClient)(nil).ListRoles), ctx, hcpHostURL)
}

// AssignGroupsToProject mocks base method
func (m *MockClient) AssignGroupsToProject(ctx context.Context, hcpHostURL, epicProjectID, projectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignGroupsToProject", ctx, hcpHostURL, epicProjectID, projectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignGroupsToProject indicates an expected call of AssignGroupsToProject
func (mr *MockClientMockRecorder) AssignGroupsToProject(ctx, hcpHostURL, epicProjectID, projectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignGroupsToProject", reflect.TypeOf((*MockClient)(nil).AssignGroupsToProject), ctx, hcpHostURL, epicProjectID, projectID)
}

// UpdateSiteAdminClaimInHCP mocks base method
func (m *MockClient) UpdateSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSiteAdminClaimInHCP", ctx, tenantID, applianceID, hcpHostURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSiteAdminClaimInHCP indicates an expected call of UpdateSiteAdminClaimInHCP
func (mr *MockClientMockRecorder) UpdateSiteAdminClaimInHCP(ctx, tenantID, applianceID, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSiteAdminClaimInHCP", reflect.TypeOf((*MockClient)(nil).UpdateSiteAdminClaimInHCP), ctx, tenantID, applianceID, hcpHostURL)
}

// RemoveSiteAdminClaimInHCP mocks base method
func (m *MockClient) RemoveSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSiteAdminClaimInHCP", ctx, tenantID, applianceID, hcpHostURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSiteAdminClaimInHCP indicates an expected call of RemoveSiteAdminClaimInHCP
func (mr *MockClientMockRecorder) RemoveSiteAdminClaimInHCP(ctx, tenantID, applianceID, hcpHostURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSiteAdminClaimInHCP", reflect.TypeOf((*MockClient)(nil).RemoveSiteAdminClaimInHCP), ctx, tenantID, applianceID, hcpHostURL)
}