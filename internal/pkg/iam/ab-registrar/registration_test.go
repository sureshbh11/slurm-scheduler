// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
package abregistrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	samlHttp "github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/iam"
	abClient "github.com/hpe-hcss/iam-lib/pkg/ab-registration-client"
	"github.com/hpe-hcss/iam-lib/pkg/resource"
)

const (
	dummyError     = "dummy error"
	errNilMsg      = "expected error is nil"
	errNotMatchMsg = "error expectation does not match"
)

var (
	mockIdentityURL   = "mockIdentityURL"
	dummyAuthBrokerID = "dummy-ab-id"
	testContext       = context.Background()
	dummyAB           = resource.AuthorizationBroker{
		ID:  dummyAuthBrokerID,
		URL: mockIdentityURL,
	}
	httpClient = &http.Client{}
)

func TestRegisterAuthorizationBroker(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func(*MockRegistrarClient, context.Context, string, string, string, string)
		expectedError error
	}{
		{
			name: "success",
			prepare: func(
				mockRegClient *MockRegistrarClient,
				ctx context.Context,
				tenantID string,
				token string,
				spaceID string,
				externalURL string,
			) {
				ab := resource.AuthorizationBroker{
					ID:          AuthorizationBrokerID,
					URL:         externalURL,
					Permissions: perms,
				}
				c := abClient.ClientDetails{
					TenantID: tenantID,
					Token:    token,
				}
				mockRegClient.EXPECT().Register(ctx, c, ab, spaceID).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "success",
			prepare: func(
				mockRegClient *MockRegistrarClient,
				ctx context.Context,
				tenantID string,
				token string,
				spaceID string,
				externalURL string,
			) {
				ab := resource.AuthorizationBroker{
					ID:          AuthorizationBrokerID,
					URL:         externalURL,
					Permissions: perms,
				}
				c := abClient.ClientDetails{
					TenantID: tenantID,
					Token:    token,
				}
				mockRegClient.EXPECT().Register(ctx, c, ab, spaceID).Return(fmt.Errorf(dummyError))
			},
			expectedError: fmt.Errorf(dummyError),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tenantID := "dummyTenantID"
			token := "dummyToken"
			spaceID := "dummySpaceID"
			externalURL := "http://localhost"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegClient := NewMockRegistrarClient(ctrl)
			tt.prepare(mockRegClient, ctx, tenantID, token, spaceID, externalURL)

			abReg := NewRegistrar(mockIdentityURL, mockRegClient, nil)
			err := abReg.RegisterAuthorizationBroker(ctx, tenantID, token, spaceID, externalURL)
			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError.Error(), err.Error(), errNilMsg)
			} else {
				assert.Equal(t, tt.expectedError, err, errNotMatchMsg)
			}
		})
	}
}

func TestAddRoles(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func(*MockRegistrarClient, context.Context, string, string, string)
		expectedError error
	}{
		{
			name: "success",
			prepare: func(
				mockRegClient *MockRegistrarClient,
				ctx context.Context,
				tenantID string,
				token string,
				spaceID string,
			) {

				c := abClient.ClientDetails{
					TenantID: tenantID,
					Token:    token,
				}
				mockRegClient.EXPECT().AddRole(
					ctx, c, AuthorizationBrokerID, gomock.Any(), spaceID,
				).AnyTimes().Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func(
				mockRegClient *MockRegistrarClient,
				ctx context.Context,
				tenantID string,
				token string,
				spaceID string,
			) {

				c := abClient.ClientDetails{
					TenantID: tenantID,
					Token:    token,
				}
				mockRegClient.EXPECT().AddRole(
					ctx, c, AuthorizationBrokerID, gomock.Any(), spaceID,
				).AnyTimes().Return(fmt.Errorf(dummyError))
			},
			expectedError: fmt.Errorf(dummyError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tenantID := "dummyTenantID"
			token := "dummyToken"
			spaceID := "dummySpaceID"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegClient := NewMockRegistrarClient(ctrl)
			tt.prepare(mockRegClient, ctx, tenantID, token, spaceID)

			abReg := NewRegistrar(mockIdentityURL, mockRegClient, nil)
			err := abReg.AddRoles(ctx, tenantID, token, spaceID)
			if tt.expectedError != nil {
				assert.NotNil(t, err.Error(), errNilMsg)
			} else {
				assert.Equal(t, tt.expectedError, err, errNotMatchMsg)
			}
		})
	}
}

func assertURL(t *testing.T, expected string, actual *url.URL) {
	expectedURL, err := url.Parse(expected)
	require.NoError(t, err)
	assert.Equal(t, expectedURL.String(), actual.String())
}

func Test_DeRegisterAuthorizationBroker(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			assertURL(t, registrationPath+"/"+dummyAB.ID, req.URL)
			http.Error(rw, "Deleted", http.StatusNoContent)
			return
		}
		http.Error(rw, "Bad request", http.StatusBadRequest)
	}))
	defer server.Close()
	tenantID := "dummy-tenant"
	token := "dummy-token"
	spaceID := "dummy-space"
	testContext = context.Background()

	mockIAMURL := "http://localhost:8080"
	mockRegClient := NewMockRegistrarClient(ctrl)
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	abReg := NewRegistrar(mockIAMURL, mockRegClient, mockHttpclient)
	mockAppsResponse := map[string]string{}
	mockHTTPResp, _ := iam.NewJSONResponse(200, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	err := abReg.DeRegisterAuthorizationBroker(testContext, tenantID, token, spaceID)
	assert.NoError(t, err)

	mockHTTPResp, _ = iam.NewJSONResponse(204, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	err = abReg.DeRegisterAuthorizationBroker(testContext, tenantID, token, spaceID)
	assert.NoError(t, err)

	errResp := "Authorization Broker " + AuthorizationBrokerID + " not found"
	mockHTTPResp, _ = iam.NewJSONResponse(404, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	err = abReg.DeRegisterAuthorizationBroker(testContext, tenantID, token, spaceID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errResp)

	errResp = "failed to delete " + AuthorizationBrokerID + " Authorization broker"
	mockErr := errors.New("Failed to execute HTTP Requeest")
	mockHTTPResp, _ = iam.NewJSONResponse(500, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = abReg.DeRegisterAuthorizationBroker(testContext, tenantID, token, spaceID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
	assert.Contains(t, err.Error(), errResp)
}

func Test_DeleteRoles(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			assertURL(t, registrationPath+"/"+dummyAB.ID, req.URL)
			http.Error(rw, "Deleted", http.StatusNoContent)
			return
		}
		http.Error(rw, "Bad request", http.StatusBadRequest)
	}))
	defer server.Close()
	tenantID := "dummy-tenant"
	token := "dummy-token"
	testContext = context.Background()

	mockIAMURL := "http://localhost:8080"
	mockRegClient := NewMockRegistrarClient(ctrl)
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	mockClientDetails := abClient.ClientDetails{
		TenantID: tenantID,
		Token:    token,
	}
	abReg := NewRegistrar(mockIAMURL, mockRegClient, mockHttpclient)
	externalRoles := []resource.RoleExternal{{URI: "/roles/role1",
		Role: resource.Role{ID: "role1", Name: "HPCaas Datacenter Admin"}},
		{URI: "/roles/role2",
			Role: resource.Role{ID: "role2", Name: "HPCaas Wrong Role"}},
		{URI: "/roles/rold3",
			Role: resource.Role{ID: "role-3", Name: "other-role-1"}},
		{URI: "/roles/role4",
			Role: resource.Role{ID: "rold-4", Name: "other-role-2"}}}
	mockRolesResponse := resource.RolesListResponse{Members: externalRoles}

	mockHTTPResp, _ := iam.NewJSONResponse(200, mockRolesResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	fmt.Println(mockHTTPResp, mockRegClient)
	mockRegClient.EXPECT().DeleteRole(gomock.Any(), mockClientDetails, "role1").Return(nil).Times(1)
	err := abReg.DeleteRoles(testContext, tenantID, token)
	assert.NoError(t, err)
}
