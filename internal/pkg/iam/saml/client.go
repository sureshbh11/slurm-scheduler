// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

//go:generate mockgen -source ./client.go -package saml -destination ./client_mock.go

// Package saml - saml package for IAM operations
package saml

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hpe-hcss/loglib/pkg/errors"
	"github.com/hpe-hcss/mlops/internal/pkg/common"
	samlHttp "github.com/hpe-hcss/mlops/internal/pkg/http"
	abregistrar "github.com/hpe-hcss/mlops/internal/pkg/iam/ab-registrar"
)

const (
	samlAppURL          = "/1.0/applications/saml/"
	appAssignmentURL    = "/1.0/app-assignment-rules"
	claimInjectionURL   = "/1.0/claim-injections"
	errorMarshalRequest = "while marshalling request body for %s"
	errorDecodeResponse = "while decoding response body"
)

//Client saml client interface
type Client interface {
	CreateSAMLApp(ctx context.Context, token string, app *Application) error
	GetSAMLAppMetadata(ctx context.Context, token string, app Application) (string, error)
	CreateSAMLAppAssigner(ctx context.Context, token string, app Application, resourceIDs, permissions []string) (string, error)
	UpdateSAMLAppAssigner(ctx context.Context, token, assignerID string, resourceIDs []string) error

	CreateSAMLClaimMappings(ctx context.Context, token string, app Application, claims []ClaimMappings) (string, error)
	UpdateSAMLClaimMappings(ctx context.Context, token, claimID string, claims []ClaimMappings) error
	DeleteSAMLAppAssigner(ctx context.Context, token, assignerID string) error
	DeleteSAMLClaimMappings(ctx context.Context, token, claimID string) error
	DeleteSAMLApp(ctx context.Context, token, samlAppID string) error
}

//ClientImpl implements Client interface for 1.0 api version
type ClientImpl struct {
	client common.HTTPClient
	iamURL string
}

//NewSAMLClient returns client interface implementation
func NewSAMLClient(iamURL string, client common.HTTPClient) *ClientImpl {
	return &ClientImpl{
		iamURL: iamURL,
		client: client,
	}
}

//CreateSAMLApp creates saml app in iam
func (c *ClientImpl) CreateSAMLApp(ctx context.Context, token string, app *Application) error {
	reqBody, err := json.Marshal(createApplicationReq{
		Name:    app.Name,
		SpaceID: app.SpaceID,
		Settings: settings{
			SignOn: signOn{
				DefaultRelayState:     "",
				SSOAcsURL:             app.SSOAcsURL,
				Recipient:             app.SSOAcsURL,
				Destination:           app.SSOAcsURL,
				SubjectNameIDFormat:   "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
				SubjectNameIDTemplate: "${user.userName}",
				ResponseSigned:        true,
				AssertionSigned:       true,
				HonorForceAuthn:       true,
				Audience:              "HPE MLOps",
				SignatureAlgorithm:    "RSA_SHA256",
				DigestAlgorithm:       "SHA256",
				AuthContextClassRef:   "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
				InlineHooks:           []string{},
				AttributeStatements:   []string{},
			},
		},
	})
	if err != nil {
		return errors.Wrapf(err, errorMarshalRequest, app.Name)
	}
	resp, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+"/1.0/applications/saml",
		http.MethodPost,
		map[string]string{"Authorization": token},
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml app create request for %s", app.Name)
	}

	var appRes createApplicationRes
	err = json.NewDecoder(resp.Body).Decode(&appRes)
	if err != nil {
		return errors.Wrap(err, errorDecodeResponse)
	}
	app.ID = appRes.ID
	app.Kid = appRes.Credentials.Signing.Kid

	return nil
}

//GetSAMLAppMetadata gets saml app metadata
func (c *ClientImpl) GetSAMLAppMetadata(ctx context.Context, token string, app Application) (string, error) {
	resp, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+"/1.0/applications/saml/"+app.ID+"/metadata/"+app.Kid,
		http.MethodGet,
		map[string]string{"Authorization": token},
		nil,
	)
	if err != nil {
		return "", errors.Wrapf(err, "while executing get saml app metadata request for %s", app.ID)
	}

	metadata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get saml metadata for appliance: %s", app.ID)
	}

	return string(metadata), nil

}

//CreateSAMLAppAssigner creates saml app assigner
func (c *ClientImpl) CreateSAMLAppAssigner(
	ctx context.Context,
	token string,
	app Application,
	resourceIDs []string,
	permissions []string,
) (string, error) {
	appAssigner := createAppAssigner{
		Name:          app.Name,
		ApplicationID: app.ID,
		SpaceID:       app.SpaceID,
		Permissions:   permissions,
	}
	var resources []resource
	for _, id := range resourceIDs {
		resources = append(resources, resource{
			ID:   id,
			ABID: abregistrar.AuthorizationBrokerID,
		})
	}
	appAssigner.Resources = resources
	reqBody, err := json.Marshal(appAssigner)
	if err != nil {
		return "", errors.Wrapf(err, errorMarshalRequest, app.Name)
	}

	resp, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+appAssignmentURL,
		http.MethodPost,
		map[string]string{"Authorization": token},
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", errors.Wrapf(err, "while executing  saml app assigner create request for %s", app.Name)
	}

	var appRes responseID
	err = json.NewDecoder(resp.Body).Decode(&appRes)
	if err != nil {
		return "", errors.Wrap(err, errorDecodeResponse)
	}
	return appRes.ID, nil
}

//UpdateSAMLAppAssigner updates saml app assigner
func (c *ClientImpl) UpdateSAMLAppAssigner(ctx context.Context, token, assignerID string, resourceIDs []string) error {
	var resources []resource
	for _, id := range resourceIDs {
		resources = append(resources, resource{
			ID:   id,
			ABID: abregistrar.AuthorizationBrokerID,
		})
	}
	reqBody, err := json.Marshal([]patchReq{{
		Op:    "replace",
		Path:  "/resources",
		Value: resources,
	}})
	if err != nil {
		return errors.Wrapf(err, errorMarshalRequest, assignerID)
	}

	_, err = samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+appAssignmentURL+"/"+assignerID,
		http.MethodPatch,
		map[string]string{"Authorization": token},
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml assigner update request for %s", assignerID)
	}

	return nil
}

//CreateSAMLClaimMappings creates saml claim mappings
func (c *ClientImpl) CreateSAMLClaimMappings(
	ctx context.Context,
	token string,
	app Application,
	claims []ClaimMappings,
) (string, error) {

	var claimValues []claimValue

	for _, claim := range claims {
		var allOfAttr []allOf
		for _, perm := range claim.Permissions {
			allOfAttr = append(allOfAttr, allOf{
				Match: match{
					Resource: resource{
						ID:   claim.ResourceID,
						ABID: abregistrar.AuthorizationBrokerID,
					},
					Permission: perm,
				},
			})
		}
		claimValues = append(claimValues, claimValue{
			Value: claim.ClaimValue,
			When: when{
				AllOf: allOfAttr,
			},
		})
	}
	reqBody, err := json.Marshal(createAppClaimMapping{
		Name:          app.Name,
		ApplicationID: app.ID,
		SpaceID:       app.SpaceID,
		ClaimMapping: []claimMapping{{
			Name:        "memberOf",
			MatchType:   "union",
			ClaimValues: claimValues,
		}},
	})
	if err != nil {
		return "", errors.Wrapf(err, errorMarshalRequest, app.Name)
	}
	resp, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+claimInjectionURL,
		http.MethodPost,
		map[string]string{"Authorization": token},
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", errors.Wrapf(err, "while executing saml claim create request for %s", app.Name)
	}

	var appRes responseID
	err = json.NewDecoder(resp.Body).Decode(&appRes)
	if err != nil {
		return "", errors.Wrap(err, errorDecodeResponse)
	}

	return appRes.ID, nil
}

//UpdateSAMLClaimMappings updates saml claim maapings
func (c *ClientImpl) UpdateSAMLClaimMappings(ctx context.Context, token, claimID string, claims []ClaimMappings) error {
	var claimValues []claimValue

	for _, claim := range claims {
		var allOfAttr []allOf
		for _, perm := range claim.Permissions {
			allOfAttr = append(allOfAttr, allOf{
				Match: match{
					Resource: resource{
						ID:   claim.ResourceID,
						ABID: abregistrar.AuthorizationBrokerID,
					},
					Permission: perm,
				},
			})
		}
		claimValues = append(claimValues, claimValue{
			Value: claim.ClaimValue,
			When: when{
				AllOf: allOfAttr,
			},
		})
	}
	reqBody, err := json.Marshal([]patchReq{{
		Op:   "replace",
		Path: "/claimMapping",
		Value: []claimMapping{{
			Name:        "memberOf",
			MatchType:   "union",
			ClaimValues: claimValues,
		}},
	}})
	if err != nil {
		return errors.Wrapf(err, errorMarshalRequest, claimID)
	}

	_, err = samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+claimInjectionURL+"/"+claimID,
		http.MethodPatch,
		map[string]string{"Authorization": token},
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml claim update request for %s", claimID)
	}
	return nil
}

// DeleteSAMLClaimMappings deletes the given saml app claim mappings resource
func (c *ClientImpl) DeleteSAMLClaimMappings(ctx context.Context, token, claimID string) error {
	_, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+claimInjectionURL+"/"+claimID,
		http.MethodDelete,
		map[string]string{"Authorization": token},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml claim delete request for %s", claimID)
	}
	return nil
}

// DeleteSAMLAppAssigner deletes the given saml app assigner resource
func (c *ClientImpl) DeleteSAMLAppAssigner(ctx context.Context, token, assignerID string) error {
	_, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+appAssignmentURL+"/"+assignerID,
		http.MethodDelete,
		map[string]string{"Authorization": token},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml assigner delete request for %s", assignerID)
	}
	return nil
}

// DeleteSAMLApp deletes the given saml app resource
func (c *ClientImpl) DeleteSAMLApp(ctx context.Context, token, samlAppID string) error {
	// First the SAML App has to be deactivated
	_, err := samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+samlAppURL+samlAppID+"/lifecycle/deactivate",
		http.MethodPost,
		map[string]string{"Authorization": token},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while deactivating saml app during delete request for %s", samlAppID)
	}

	// Now, delete the SAML App
	_, err = samlHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		c.iamURL+samlAppURL+samlAppID,
		http.MethodDelete,
		map[string]string{"Authorization": token},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while executing saml app delete request for %s", samlAppID)
	}
	return nil
}
