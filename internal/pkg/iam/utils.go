// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	hpeErrors "github.com/hpe-hcss/errors/pkg/errors"

	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
)

const appLinksTemplate = "scim/v1/extensions/tenant/%s/Users/%s/appLinks"

// HTTPResponse contains AppLinks Response object
type HTTPResponse struct {
	Label  string `json:"label"`
	SSOURL string `json:"linkUrl"`
}

// NewJSONResponse returns sample HTTP JSON response
func NewJSONResponse(status int, body interface{}) (*http.Response, error) {
	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp := &http.Response{
		Status:        strconv.Itoa(status),
		StatusCode:    status,
		Body:          ioutil.NopCloser(bytes.NewReader(encoded)),
		Header:        http.Header{},
		ContentLength: -1,
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp, nil
}

// GetSSOURL calls the Identity api which in turns calls Okta to find the
// SSO URL corresponding to a provisioned service instance
func GetSSOURL(iamURL, appLabel, tenantID, userID string, authToken string, httpClient common.HTTPClient) (HTTPResponse, error) {

	appLinksURI := fmt.Sprintf(appLinksTemplate, tenantID, userID)
	url := fmt.Sprintf("%v/%v", iamURL, appLinksURI)
	log.Info(appLinksURI)
	log.Info(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("Failed to get SSO URL! Error: %v\n", err)
		return HTTPResponse{}, getInternalServerError()
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("Failed to get SSO URL! Error: %v\n", err)
		return HTTPResponse{}, getInternalServerError()
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read the response body. Error: %v", err)
		return HTTPResponse{}, getInternalServerError()
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf("Failed to get SSO URL from IAM! Error: %v", resp.StatusCode)
		return HTTPResponse{}, getInternalServerError()
	}

	var output []HTTPResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Errorf("Failed to parse applinks response from IAM! Error: %v\n", err)
		return HTTPResponse{}, getInternalServerError()
	}

	if len(output) == 0 {
		log.Errorf("Unable to find members in appLinks response")
		errorResponse := hpeErrors.ErrorResponse{
			Message: fmt.Sprintf("SSO URL not found for the given Tenant %s", tenantID),
		}
		return HTTPResponse{}, hpeErrors.MakeErrNotFound(errorResponse)
	}

	// walk through the list and find the member that matches
	for _, entry := range output {
		if strings.ToLower(entry.Label) == appLabel {
			log.Infof("Found matching SSO URL! %s", entry.Label)
			return entry, nil
		}
	}

	// SAML App not found
	log.Errorf("No App name found for the given Tenant having label %s", appLabel)
	errorResponse := hpeErrors.ErrorResponse{
		Message: fmt.Sprintf("SSO URL not found for the given Tenant %s", tenantID),
	}
	return HTTPResponse{}, hpeErrors.MakeErrNotFound(errorResponse)
}

// getInternalServerError returns custom InternalServer Error object
func getInternalServerError() error {
	errorResponse := hpeErrors.ErrorResponse{
		Message:   "Failed to get SSO URL from IAM",
		ErrorCode: "Internal Server Error",
	}
	return hpeErrors.MakeErrInternalError(errorResponse)
}
