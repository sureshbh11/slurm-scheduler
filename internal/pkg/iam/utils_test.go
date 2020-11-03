// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
)

const (
	mockIamURL      = "http://mockUrl:8080"
	mockLabelPrefix = "hpe-HPCaas-hcp-"
	mockTenantID    = "mockTenant"
	mockUserID      = "mockUser"
	mockToken       = "mockToken"
)

func Test_GetSSOURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	appLinksURI := fmt.Sprintf(appLinksTemplate, mockTenantID, mockUserID)
	mockIAMURL := fmt.Sprintf("%v/%v", mockIamURL, appLinksURI)
	mockHttpclient := common.NewMockHTTPClient(ctrl)
	mockAppsResponse := []HTTPResponse{}
	mockAppsResponse = append(mockAppsResponse,
		HTTPResponse{Label: "hpe-HPCaas-hcp-mockInstance",
			SSOURL: mockIamURL})
	errMsg := "Failed to get SSO URL from IAM"
	successResp, _ := NewJSONResponse(200, mockAppsResponse)
	errResp, _ := NewJSONResponse(500, mockAppsResponse)
	emptyResp, _ := NewJSONResponse(200, []HTTPResponse{})
	wrongResp, _ := NewJSONResponse(200, "Wrong Reponse")
	mockReq, err := http.NewRequest(http.MethodGet, mockIAMURL, nil)
	mockReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockToken))

	// Http Response returns 500
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(errResp, nil)
	ssoResponse, err := GetSSOURL(mockIamURL, mockLabelPrefix, mockTenantID, mockUserID, mockToken, mockHttpclient)
	assert.Equal(t, "", ssoResponse.SSOURL)
	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())

	// Http Response returns Error
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("Failed"))
	ssoResponse, err = GetSSOURL(mockIamURL, mockLabelPrefix, mockTenantID, mockUserID, mockToken, mockHttpclient)
	assert.Equal(t, "", ssoResponse.SSOURL)
	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())

	// Http Response Parsing Error
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(wrongResp, nil)
	ssoResponse, err = GetSSOURL(mockIamURL, mockLabelPrefix, mockTenantID, mockUserID, mockToken, mockHttpclient)
	assert.Equal(t, "", ssoResponse.SSOURL)
	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())

	// Http Response returns empty list
	errMsg = "SSO URL not found for the given Tenant " + mockTenantID
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(emptyResp, nil)
	ssoResponse, err = GetSSOURL(mockIamURL, mockLabelPrefix, mockTenantID, mockUserID, mockToken, mockHttpclient)
	assert.Equal(t, "", ssoResponse.SSOURL)
	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())

	// Http Response not having requried entry
	mockAppsResponse[0].Label = "Wrong Label"
	successResp, _ = NewJSONResponse(200, mockAppsResponse)
	mockHttpclient.EXPECT().Do(gomock.Any()).Return(successResp, nil)
	ssoResponse, err = GetSSOURL(mockIamURL, mockLabelPrefix, mockTenantID, mockUserID, mockToken, mockHttpclient)
	assert.Equal(t, "", ssoResponse.SSOURL)
	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())
}
