//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package abregistrar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	hpeErrors "github.com/hpe-hcss/errors/pkg/errors"
	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/common"
	abClient "github.com/hpe-hcss/iam-lib/pkg/ab-registration-client"
	"github.com/hpe-hcss/iam-lib/pkg/resource"
	"github.com/hpe-hcss/loglib/pkg/errors"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
)

// AuthorizationBrokerID is the public broker id for IAM broker
const AuthorizationBrokerID = "HPCaas"

var (
	perms               []resource.Permission
	jsonContent         = "application/json"
	contentType         = "Content-Type"
	bearer              = "Bearer %s"
	authorizationHeader = "Authorization"
	registrationPath    = "/v1alpha2/authorization-brokers"
	rolePath            = "/v1alpha2/roles"
)

func init() {
	for _, v := range permissions {
		perms = append(perms, v)
	}
}

// ABRegistration implements Registrar interface
type ABRegistration struct {
	iamURL     string
	client     RegistrarClient
	httpClient common.HTTPClient
}

// RegisterAuthorizationBroker calls the IAM service and passes in the HPCaas permissions
func (abReg *ABRegistration) RegisterAuthorizationBroker(ctx context.Context,
	tenantID, token, spaceID, externalURL string) error {

	ab := resource.AuthorizationBroker{
		ID:          AuthorizationBrokerID,
		URL:         externalURL,
		Permissions: perms,
	}

	c := abClient.ClientDetails{
		TenantID: tenantID,
		Token:    token,
	}
	err := abReg.client.Register(ctx, c, ab, spaceID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// AddRoles calls the IAM service and passes in the Azure roles
func (abReg *ABRegistration) AddRoles(ctx context.Context,
	tenantID, token, spaceID string) error {

	c := abClient.ClientDetails{
		TenantID: tenantID,
		Token:    token,
	}

	for _, role := range roles {
		err := abReg.client.AddRole(ctx, c, AuthorizationBrokerID, role, spaceID)
		if err != nil {
			return errors.Wrapf(err, "while adding role: %v, spaceID: %s", role, spaceID)
		}
	}

	return nil
}

// DeleteRoles iteratively deletes all the roles pushed by HPCaas service to the tenant
func (abReg *ABRegistration) DeleteRoles(ctx context.Context,
	tenantID, token string) error {

	var errMsg string
	c := abClient.ClientDetails{
		TenantID: tenantID,
		Token:    token,
	}

	log.Infof("Fetching list of roles for tenant %s", tenantID)
	roles, err := abReg.getExistingRolesWithID(ctx, token)
	if err != nil {
		log.Errorf("failed in fetching list of roles for tenant %s: %v", tenantID, err)
		return err
	}

	log.Infof("Fetched total of %d HPCaas roles for tenant %s", len(roles), tenantID)
	for i, role := range roles {
		log.Infof("Deleting %s role", role.Name)
		err := abReg.client.DeleteRole(ctx, c, role.ID)

		if err != nil {
			log.Errorf("Failed to delete roles for the tenant %s", tenantID)
			errMsg = errMsg + fmt.Sprintf("(%d)failed to delete role %s with err: %v. ", i, role.Name, err)
		} else {
			log.Infof("Successfully Deleted %s from tenant %s", role.Name, tenantID)
		}
	}

	if errMsg != "" {
		return fmt.Errorf(errMsg)
	}
	return nil
}

// NewRegistrar constructs a ABRegistration that will register the HPCaas
// service with the IAM service
func NewRegistrar(iamURL string, c RegistrarClient, client common.HTTPClient) Registrar {
	url := strings.TrimRight(iamURL, "/")
	return &ABRegistration{
		client:     c,
		iamURL:     url,
		httpClient: client,
	}
}

// DeRegisterAuthorizationBroker is to delete the existing authorization-brokers and their permissions for the tenant
func (abReg *ABRegistration) DeRegisterAuthorizationBroker(ctx context.Context, tenantID string, token string, spaceID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Delete AB(controller)")
	defer span.Finish()
	var errMsg string
	brokerID := AuthorizationBrokerID
	url := abReg.iamURL + registrationPath + "/" + brokerID
	log.Infof("Executing delete of authz-broker %s", brokerID)

	response, err := abReg.callHTTPMethod(ctx, token, http.MethodDelete, url, nil)
	if err != nil {
		errMsg = fmt.Sprintf("failed to delete %s Authorization broker for %v tenant. Err: %v", brokerID, tenantID, err)
		return errors.Wrap(err, errMsg)
	}

	defer closeBody(response)
	if response.StatusCode == 204 {
		log.Infof("Successfully removed %s Authorization Broker", brokerID)
		return nil

	} else if response.StatusCode == 404 {
		errMsg = fmt.Sprintf("Authorization Broker %s not found: %v", brokerID, err)
		return hpeErrors.MakeErrNotFound(hpeErrors.ErrorResponse{Message: errMsg})

	}
	return nil
}

func (abReg *ABRegistration) callHTTPMethod(ctx context.Context, token, method, url string, body interface{}) (*http.Response, error) {
	var reader io.Reader
	if body != nil {
		rpBytes, err := json.Marshal(body)
		if err != nil {
			log.Errorf("could not marshal body %+v: %v", body, err)
			return nil, err
		}
		reader = bytes.NewBuffer(rpBytes)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Add(contentType, jsonContent)

	// Add the token to the request header
	req.Header.Add(authorizationHeader, fmt.Sprintf(bearer, token))

	resp, err := abReg.httpClient.Do(req)
	if err != nil {
		log.Errorf("%v request to %v failed: %v", method, url, err)
		return nil, err
	}
	return resp, nil
}

// Used to log a body close error if one occurs
func closeBody(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		log.Errorf("failed to close response body: %v", err)
	}
}

func (abReg *ABRegistration) getExistingRolesWithID(ctx context.Context, token string) ([]resource.Role, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "get role-ids")
	defer span.Finish()

	url := abReg.iamURL + rolePath

	response, err := abReg.callHTTPMethod(ctx, token, http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("failed to get all roles with error: %v", err)
		return []resource.Role{}, err
	}

	defer closeBody(response)
	if response.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Failed to retrieve role from %s, expected %v but status was %v",
			url, http.StatusOK, response.StatusCode)
		log.Error(ctx, msg)
		return []resource.Role{}, fmt.Errorf(msg)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("failed to read response body: %v", err)
		return []resource.Role{}, err
	}

	existing := resource.RolesListResponse{}
	err = json.Unmarshal(responseBody, &existing)
	if err != nil {
		log.Errorf("could not unmarshal response: %v", err)
		return []resource.Role{}, err
	}

	// Filter out only HPCaas roles
	members := existing.Members
	roles := filterHPCaasRoles(members)

	if len(roles) == 0 {
		log.Info("No roles returned for HPCaas service")
	}

	return roles, nil
}

func filterHPCaasRoles(members []resource.RoleExternal) []resource.Role {

	var HPCaasroles []resource.Role
	for _, member := range members {
		log.Info("Fetch the list of roles", members)
		for _, eachRole := range roles {
			if eachRole.Name == member.Name {
				r := resource.Role{
					ID:   member.ID,
					Name: member.Name,
				}
				HPCaasroles = append(HPCaasroles, r)
				log.Info("Filtered HPCaas roles", HPCaasroles)
				break
			}
		}
	}
	return HPCaasroles
}
