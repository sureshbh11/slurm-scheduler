// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package hcp

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	samlHttp "github.com/hpe-hcss/mlops/internal/pkg/common"
	hcpModels "github.com/hpe-hcss/mlops/internal/pkg/hcp/models"
	"github.com/hpe-hcss/mlops/internal/pkg/iam"
	"github.com/hpe-hcss/mlops/internal/pkg/models"
	"github.com/hpe-hcss/mlops/internal/pkg/patch"
)

const (
	testProjectID = "projID12345"
)

func Test_ConfigureAuthInHcp(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	epicInput := models.EpicConfigInput{}
	mockAppsResponse := map[string]string{}
	mockHTTPResp, _ := iam.NewJSONResponse(200, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	err := hcpClient.ConfigureAuthInHcp(mockContext, mockHostURL, epicInput)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.ConfigureAuthInHcp(mockContext, mockHostURL, epicInput)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.ConfigureAuthInHcp(mockContext, mockHostURL, epicInput)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.ConfigureAuthInHcp(mockContext, mockHostURL, epicInput)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_CreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	input := models.HCPCreateProjectInput{}
	mockHTTPResp, _ := iam.NewJSONResponse(201, nil)
	mockHTTPResp.Header.Add("Location", "/api/v1/project/"+testProjectID)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	projID, err := hcpClient.CreateProject(mockContext, mockHostURL, input)
	assert.Nil(t, err)
	assert.Equal(t, testProjectID, projID)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	projID, err = hcpClient.CreateProject(mockContext, mockHostURL, input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	projID, err = hcpClient.CreateProject(mockContext, mockHostURL, input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	projID, err = hcpClient.CreateProject(mockContext, mockHostURL, input)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_DeleteProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(204, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	err := hcpClient.DeleteProject(mockContext, mockHostURL, testProjectID)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.DeleteProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.DeleteProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.DeleteProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_UpdateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	patchData := patch.JSONPatch{{
		Op:    "replace",
		Path:  "/label/name",
		Value: "Test Project",
	}}

	mockHTTPResp, _ := iam.NewJSONResponse(201, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	err := hcpClient.UpdateProject(mockContext, mockHostURL, testProjectID, patchData)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.UpdateProject(mockContext, mockHostURL, testProjectID, patchData)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.UpdateProject(mockContext, mockHostURL, testProjectID, patchData)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.UpdateProject(mockContext, mockHostURL, testProjectID, patchData)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_ListProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.ListProject(mockContext, mockHostURL)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListProject(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListProject(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListProject(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_GetProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.GetProject(mockContext, mockHostURL, testProjectID)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetProject(mockContext, mockHostURL, testProjectID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_GetClusters(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.GetClusters(mockContext, mockHostURL)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetClusters(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetClusters(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetClusters(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_GetModels(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.GetModels(mockContext, mockHostURL)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetModels(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetModels(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetModels(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_ListRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.GetModels(mockContext, mockHostURL)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListRoles(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListRoles(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.ListRoles(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockResp := hcpModels.RoleListResource{}
	mockRole := hcpModels.RoleResource{Links: hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}},
		Label: hcpModels.CommonLabel{Name: "MockRole", Description: "Unit Testing"}}
	mockResp.Links = hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}}
	mockResp.Embedded = hcpModels.RoleListEmbeddedResource{Roles: []hcpModels.RoleResource{mockRole}}
	mockBody, _ := iam.NewJSONResponse(200, mockResp)
	getSession := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1).After(getSession)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	roles, err := hcpClient.ListRoles(mockContext, mockHostURL)
	assert.Nil(t, err)
	assert.NotNil(t, roles)
	assert.Equal(t, 1, len(roles))
	assert.Equal(t, "MockRole", roles[0].Label.Name)
	assert.Equal(t, "Unit Testing", roles[0].Label.Description)

	mockBody, _ = iam.NewJSONResponse(404, mockResp)
	getSession = mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1).After(getSession)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	roles, err = hcpClient.ListRoles(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "status code: 404")
	assert.Equal(t, []hcpModels.RoleResource{}, roles)

	mockBody, _ = iam.NewJSONResponse(200, "wrong response")
	getSession = mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1).After(getSession)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	roles, err = hcpClient.ListRoles(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Failed to unmarshal body")
}

func Test_AssignGroupsToProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockProjectID := "mockProjectID"
	mockSiteID := "mockSiteID"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err := hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "while creating session")
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockResp := hcpModels.RoleListResource{}
	mockAdminRole := hcpModels.RoleResource{Links: hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}},
		Label: hcpModels.CommonLabel{Name: "Admin", Description: "Mock Admin Role"}}
	mockMemberRole := hcpModels.RoleResource{Links: hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/2"}},
		Label: hcpModels.CommonLabel{Name: "Member", Description: "Mock Member Role"}}
	mockResp.Embedded = hcpModels.RoleListEmbeddedResource{Roles: []hcpModels.RoleResource{mockAdminRole}}
	mockBody, _ := iam.NewJSONResponse(200, mockResp)

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockBody, _ = iam.NewJSONResponse(204, mockResp)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Failed to get admin and or member role(s)")

	mockResp.Embedded.Roles = append(mockResp.Embedded.Roles, mockMemberRole)
	mockBody, _ = iam.NewJSONResponse(200, mockResp)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockBody, _ = iam.NewJSONResponse(204, mockResp)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.Nil(t, err)

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockBody, _ = iam.NewJSONResponse(204, mockResp)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, mockErr).Times(1)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Failed to Add Groups")
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockBody, _ = iam.NewJSONResponse(204, mockResp)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, mockErr).Times(1)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "while deleting the session")
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockBody, _ = iam.NewJSONResponse(200, mockResp)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1)
	err = hcpClient.AssignGroupsToProject(mockContext, mockHostURL, mockProjectID, mockSiteID)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "incorrect status code")
}

func Test_GetConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)

	mockHTTPResp, _ := iam.NewJSONResponse(200, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(3)
	_, err := hcpClient.GetConfig(mockContext, mockHostURL)
	assert.Nil(t, err)

	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetConfig(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetConfig(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, mockErr).Times(1)
	_, err = hcpClient.GetConfig(mockContext, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())
}

func Test_UpdateSiteAdminClaimInHCP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	testTenant := "testTenant"
	mockApplianceID := "mockAppId"
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)
	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)

	// Get Session failure case
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(nil, mockErr).Times(1)
	err := hcpClient.UpdateSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	// Site Admin Update failure: List Project failure case
	mockAppsResponse := map[string]string{}
	mockHTTPResp, _ := iam.NewJSONResponse(200, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(nil, mockErr).Times(1)
	err = hcpClient.UpdateSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	// Site Admin Update failure: Project ref not present in List Project response case
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	err = hcpClient.UpdateSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "while getting site admin project ref from Hcp")

	// Site Admin Update success: List Role success case
	tenants := hcpModels.ListTenants{}
	tenants.Links.Self.Href = "/v1/mockTenant/1"
	mockTenant := hcpModels.Tenant{}
	mockTenant.Links.Self.Href = "/v1/mockTenant/1"
	mockTenant.Label.Name = siteAdminProject
	//	mockTenant.Features.MlProject = true
	tenants.Embedded = hcpModels.Embedded{Tenants: []hcpModels.Tenant{mockTenant}}
	mockBody, _ := iam.NewJSONResponse(200, tenants)

	roleMockResp := hcpModels.RoleListResource{}
	mockRole := hcpModels.RoleResource{Links: hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}},
		Label: hcpModels.CommonLabel{Name: "Site Admin", Description: "Unit Testing"}}
	roleMockResp.Links = hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}}
	roleMockResp.Embedded = hcpModels.RoleListEmbeddedResource{Roles: []hcpModels.RoleResource{mockRole}}
	roleMockBody, _ := iam.NewJSONResponse(200, roleMockResp)

	firstCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	secondCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1).After(firstCall)
	thirdCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1).After(secondCall)
	fourthCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(roleMockBody, nil).Times(1).After(thirdCall)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1).After(fourthCall)
	mockNoContentResp, _ := iam.NewJSONResponse(204, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockNoContentResp, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil)
	err = hcpClient.UpdateSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.Nil(t, err)

}

func Test_RemoveSiteAdminClaimInHCP(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	testTenant := "testTenant"
	mockApplianceID := "mockAppId"
	mockHostURL := "http://mockHcpURL:8080"
	mockHttpclient := samlHttp.NewMockHTTPClient(ctrl)
	hcpClient := NewHCPClient(mockHttpclient)
	errMsg := "Failed to complete request"
	mockErr := errors.New(errMsg)

	// Get Session failure case
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(nil, mockErr).Times(1)
	err := hcpClient.RemoveSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	// Site Admin claim deletion failure: List Project failure case
	mockAppsResponse := map[string]string{}
	mockHTTPResp, _ := iam.NewJSONResponse(200, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(nil, mockErr).Times(1)
	err = hcpClient.RemoveSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), mockErr.Error())

	// Site Admin claim deletion failure: Project ref not present in List Project response case
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(2)
	err = hcpClient.RemoveSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "while getting site admin project ref from Hcp")

	// Site Admin Update success: List Role success case
	tenants := hcpModels.ListTenants{}
	tenants.Links.Self.Href = "/v1/mockTenant/1"
	mockTenant := hcpModels.Tenant{}
	mockTenant.Links.Self.Href = "/v1/mockTenant/1"
	mockTenant.Label.Name = siteAdminProject
	//	mockTenant.Features.MlProject = true
	tenants.Embedded = hcpModels.Embedded{Tenants: []hcpModels.Tenant{mockTenant}}
	mockBody, _ := iam.NewJSONResponse(200, tenants)

	roleMockResp := hcpModels.RoleListResource{}
	mockRole := hcpModels.RoleResource{Links: hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}},
		Label: hcpModels.CommonLabel{Name: "Site Admin", Description: "Unit Testing"}}
	roleMockResp.Links = hcpModels.CommonResourceLinks{Self: hcpModels.CommonLink{HRef: "/v1/mockRole/1"}}
	roleMockResp.Embedded = hcpModels.RoleListEmbeddedResource{Roles: []hcpModels.RoleResource{mockRole}}
	roleMockBody, _ := iam.NewJSONResponse(200, roleMockResp)

	firstCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1)
	secondCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockBody, nil).Times(1).After(firstCall)
	thirdCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1).After(secondCall)
	fourthCall := mockHttpclient.EXPECT().Do(gomock.Any()).Return(roleMockBody, nil).Times(1).After(thirdCall)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil).Times(1).After(fourthCall)
	mockNoContentResp, _ := iam.NewJSONResponse(204, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockNoContentResp, nil)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(mockHTTPResp, nil)
	err = hcpClient.RemoveSiteAdminClaimInHCP(mockContext, testTenant, mockApplianceID, mockHostURL)
	assert.NotNil(t, err)
}
