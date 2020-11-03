// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

//go:generate mockgen -source ./client.go -package hcp -destination ./client_mock.go

// Package hcp - hcp package contains methods to connect to HCP Appliances
package hcp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"

	"github.com/hpe-hcss/loglib/pkg/errors"
	"github.com/hpe-hcss/monitoring/pkg/metrics"

	"github.com/hpe-hcss/mlops/internal/pkg/common"
	hcpModels "github.com/hpe-hcss/mlops/internal/pkg/hcp/models"
	mlopsHttp "github.com/hpe-hcss/mlops/internal/pkg/http"
	"github.com/hpe-hcss/mlops/internal/pkg/models"
	"github.com/hpe-hcss/mlops/internal/pkg/patch"
)

const (
	externalSvcName = "mlops-controller"
	// Failure constant used to indicate metrics failure
	Failure = "failure"
	// Success constant used to indicate metrics success
	Success             = "success"
	errorMarshalRequest = "while marshalling request body for %s"
	errorDecodeResponse = "while decoding response body"
)

var (
	status      = Failure
	hcpUserName = common.GetEnv("HCP_SERVICE_USERNAME", "dummyUser")
	hcpPassword = common.GetEnv("HCP_SERVICE_PASSWORD", "dummyPassword")
)

// Client HCP client interface
type Client interface {
	ConfigureAuthInHcp(ctx context.Context, hcpHostURL string,
		input models.EpicConfigInput) error
	CreateProject(ctx context.Context, hcpHostURL string, input models.HCPCreateProjectInput) (string, error)
	DeleteProject(ctx context.Context, hcpHostURL, projectID string) error
	ListProject(ctx context.Context, hcpHostURL string) ([]hcpModels.Tenant, error)
	GetProject(ctx context.Context, hcpHostURL, projectID string) (hcpModels.Tenant, error)
	UpdateProject(ctx context.Context, hcpHostURL, projectID string, input patch.JSONPatch) error
	GetClusters(ctx context.Context, hcpHostURL string) (models.ClusterResp, error)
	GetModels(ctx context.Context, hcpHostURL string) (models.ModelResp, error)
	GetConfig(ctx context.Context, hcpHostURL string) (hcpModels.HcpConfig, error)
	ListRoles(ctx context.Context, hcpHostURL string) ([]hcpModels.RoleResource, error)
	AssignGroupsToProject(ctx context.Context, hcpHostURL, epicProjectID, projectID string) error
	UpdateSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error
	RemoveSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error
}

//ClientImpl implements Client interface for 1.0 api versions
type ClientImpl struct {
	client common.HTTPClient
}

//NewHCPClient returns client interface implementation
func NewHCPClient(client common.HTTPClient) *ClientImpl {
	return &ClientImpl{
		client: client,
	}
}

func (c *ClientImpl) getSession(ctx context.Context, hcpHostURL, username, password string) (string, error) {

	monitor := metrics.StartExternalCall(externalSvcName, "Create Session in HCP")
	status = Failure
	defer func() { monitor.RecordWithStatus(status) }()

	requestBody, _ := json.Marshal(hcpModels.SessionInput{Name: username, Password: password})
	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+sessionPath,
		http.MethodPost,
		map[string]string{"Content-Type": "application/json"},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return "", errors.Wrapf(err, "while creating session for the user %s", username)
	}

	status = Success
	if resp.StatusCode == http.StatusCreated {
		sessionID := resp.Header.Get("location")
		return sessionID, nil
	}
	return "", errors.Wrapf(err, "while creating session for the user %s", username)
}

func (c *ClientImpl) deleteSession(ctx context.Context, hcpHostURL, session string) error {

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Create Session in HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	_, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+session,
		http.MethodDelete,
		map[string]string{sessionHeader: session},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while deleting the session.")
	}

	status = Success
	return nil
}

//ConfigureAuthInHcp updates the saml config in respective HCP Controller
func (c *ClientImpl) ConfigureAuthInHcp(ctx context.Context, hcpHostURL string,
	input models.EpicConfigInput) error {
	sessionID, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return err
	}
	requestBody, _ := json.Marshal(input)
	_, err = mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+configPath,
		http.MethodPost,
		map[string]string{sessionHeader: sessionID},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return errors.Wrapf(err, "while updating the SAML config to HCP.")
	}
	err = c.deleteSession(ctx, hcpHostURL, sessionID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateSiteAdminClaimInHCP updates the site admin project user group
func (c *ClientImpl) UpdateSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Update Site Admin Claim in Hcp")
	defer span.Finish()
	// Get List of Tenants/projects from HCP. Get Tenant Id for which matches "Site Admin" as tenant Name"
	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Update SiteAdmin Claim in HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	sessionID, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return errors.Wrapf(err, "while getting session for update Site Admin User Group.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1,
		http.MethodGet,
		map[string]string{sessionHeader: sessionID},
		bytes.NewReader(nil),
	)
	if err != nil {
		return errors.Wrapf(err, "while fetching projects from MLOps controller platform.")
	}

	var projects hcpModels.ListTenants
	_, parseRespErr := common.ParseResponse(resp, &projects)
	if parseRespErr != nil {
		log.Errorf("Failed to fetch projects from HCP: %v", parseRespErr)
		return errors.Wrapf(parseRespErr, "Failed to fetch projects from HCP")
	}

	var siteAdminProjectHRef string
	for _, project := range projects.Embedded.Tenants {
		if project.Label.Name == siteAdminProject {
			siteAdminProjectHRef = project.Links.Self.Href
			break
		}
	}
	if siteAdminProjectHRef == "" {
		return fmt.Errorf("failed while getting site admin project ref from Hcp.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}
	// Get roles reference for which  role name matches "Site Admin"
	roles, err := c.ListRoles(ctx, hcpHostURL)
	if err != nil {
		return errors.Wrapf(err, "failed while listing roles from Hcp. TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}
	var siteAdminRoleHRef string
	for _, role := range roles {
		if role.Label.Name == siteAdminRole {
			siteAdminRoleHRef = role.Links.Self.HRef
			break
		}
	}

	if siteAdminRoleHRef == "" {
		return fmt.Errorf("while getting site admin role ref from Hcp.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}

	/* Adding UserGroup */
	args := hcpModels.ExternalUserGroupListResource{
		ExternalUserGroups: []hcpModels.ExternalUserGroupResource{
			{
				Role:  siteAdminRoleHRef,
				Group: siteAdminClaimValue + "-" + applianceID,
			},
		},
	}
	requestBody, _ := json.Marshal(args)
	resp, err = mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+siteAdminProjectHRef+externalUserGroupsFlag,
		http.MethodPut,
		map[string]string{sessionHeader: sessionID},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return errors.Wrapf(err, "Updating User group in Site Admin Project in HCP failed: %v for TenantID: %s ApplianceID: %s", tenantID, applianceID, err)
	}
	if code := resp.StatusCode; code != http.StatusNoContent {
		return fmt.Errorf("Unable to update User group for Site Admin project %v: %d.TenantID: %s ApplianceID: %s", siteAdminProjectHRef, resp.StatusCode, tenantID, applianceID)
	}

	status = Success
	err = c.deleteSession(ctx, hcpHostURL, sessionID)
	if err != nil {
		return errors.Wrapf(err, "while detting session for update Site Admin User Group.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}
	return nil
}

// CreateProject is wrapper function which will contact HCP to create a tenant
func (c *ClientImpl) CreateProject(ctx context.Context, hcpHostURL string, input models.HCPCreateProjectInput) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Create HCP Project")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return "", err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Create Project in HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	requestBody, _ := json.Marshal(input)
	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1,
		http.MethodPost,
		map[string]string{sessionHeader: session},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return "", errors.Wrapf(err, "while creating project in MLOps controller platform.")
	}

	status = Success

	// Extract the ID fof the project from the response headers
	projectURI := resp.Header.Get("Location")
	projectParameters := strings.Split(projectURI, "/")
	projectID := projectParameters[len(projectParameters)-1]

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return "", err
	}
	return projectID, nil
}

// DeleteProject is wrapper function which will contact HCP to delete a tenant
func (c *ClientImpl) DeleteProject(ctx context.Context, hcpHostURL, projectID string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Delete HCP Project")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Delete Project from HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	_, err = mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1+projectID,
		http.MethodDelete,
		map[string]string{sessionHeader: session},
		nil,
	)
	if err != nil {
		return errors.Wrapf(err, "while deleting project in MLOps controller platform.")
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProject is wrapper function which will contact HCP to update a tenant
func (c *ClientImpl) UpdateProject(ctx context.Context, hcpHostURL, projectID string, input patch.JSONPatch) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Update HCP Project")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Update Project in HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	requestBody, _ := json.Marshal(input)
	_, err = mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV2+projectID,
		http.MethodPatch,
		map[string]string{sessionHeader: session},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return errors.Wrapf(err, "while updating project in MLOps controller platform.")
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return err
	}
	return nil
}

// ListProject is wrapper function which will contact HCP to list of tenants
func (c *ClientImpl) ListProject(ctx context.Context, hcpHostURL string) ([]hcpModels.Tenant, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "List HCP Projects")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return nil, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "List Projects from HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "while fetching projects from MLOps controller platform.")
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return nil, err
	}

	var tenants hcpModels.ListTenants
	_, parseRespErr := common.ParseResponse(resp, &tenants)
	if parseRespErr != nil {
		log.Errorf("Failed to fetch projects from HCP: %v", parseRespErr)
		return nil, errors.Wrapf(parseRespErr, "Failed to fetch projects from HCP")
	}

	// filter only ML projects
	var mlProjects []hcpModels.Tenant
	for _, tenant := range tenants.Embedded.Tenants {
		if tenant.Features.MlProject {
			mlProjects = append(mlProjects, tenant)
		}
	}
	return mlProjects, nil
}

// GetProject is wrapper function which will contact HCP to fetch tenant info
func (c *ClientImpl) GetProject(ctx context.Context, hcpHostURL, projectID string) (hcpModels.Tenant, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Get HCP Project")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return hcpModels.Tenant{}, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Get Project from HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1+projectID,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return hcpModels.Tenant{}, errors.Wrapf(err, "while fetching project from MLOps controller platform.")
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return hcpModels.Tenant{}, err
	}

	var tenants hcpModels.Tenant
	_, parseRespErr := common.ParseResponse(resp, &tenants)
	if parseRespErr != nil {
		log.Errorf("Failed to fetch project from HCP: %v", parseRespErr)
		return hcpModels.Tenant{}, errors.Wrapf(parseRespErr, "Failed to fetch project from HCP")
	}

	return tenants, nil
}

// GetClusters is a function which is used to fetch cluster details from HCP
func (c *ClientImpl) GetClusters(ctx context.Context, hcpHostURL string) (models.ClusterResp, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Get Clusters")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return models.ClusterResp{}, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Fetch Clusters")
	defer func() { monitor.RecordWithStatus(status) }()

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+clusterPathV2,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return models.ClusterResp{}, errors.Wrapf(err, "while fetching clusters in MLOps controller platform.")
	}
	resp.Body.Close()

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return models.ClusterResp{}, err
	}

	clustersResp := models.ClusterResp{}
	json.NewDecoder(resp.Body).Decode(&clustersResp)

	return clustersResp, nil
}

// GetModels is a function which is used to fetch model details from HCP
func (c *ClientImpl) GetModels(ctx context.Context, hcpHostURL string) (models.ModelResp, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Get Models")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return models.ModelResp{}, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Fetch Models")
	defer func() { monitor.RecordWithStatus(status) }()

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+modelPathV2,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return models.ModelResp{}, errors.Wrapf(err, "while fetching models in MLOps controller platform.")
	}
	resp.Body.Close()

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return models.ModelResp{}, err
	}

	modelsResp := models.ModelResp{}
	json.NewDecoder(resp.Body).Decode(&modelsResp)

	return modelsResp, nil
}

// ListRoles returns the list of pre-defined roles in the given HCP Controller.
func (c *ClientImpl) ListRoles(ctx context.Context, hcpHostURL string) ([]hcpModels.RoleResource, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "List Roles")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return []hcpModels.RoleResource{}, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "List Roles in HCP")
	defer func() { monitor.RecordWithStatus(status) }()
	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+rolePathV1,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return []hcpModels.RoleResource{}, errors.Wrapf(err, "ListRoles: Failed to get roles from HCP: %v", err)
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ListRoles: Failed to decode response body: %v", err)
	}
	var roleOutput hcpModels.RoleListResource
	if err := json.Unmarshal(body, &roleOutput); err != nil {
		return nil, fmt.Errorf("ListRoles: Failed to unmarshal body: %v", err)
	}

	return roleOutput.Embedded.Roles, nil
}

// AssignGroupsToProject assign groups to a Project
func (c *ClientImpl) AssignGroupsToProject(ctx context.Context, hcpHostURL, epicProjectID, projectID string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "Assign Groups to Project")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Update External Groups in HCP")
	defer func() { monitor.RecordWithStatus(status) }()
	// get existing roles
	existingRoles, err := c.ListRoles(ctx, hcpHostURL)
	if err != nil {
		return fmt.Errorf("AssignGroupsToProject: Failed to get available roles: %v", err)
	}

	var adminRoleHref string
	var memberRoleHref string

	for _, eachRole := range existingRoles {
		if eachRole.Label.Name == adminRole {
			adminRoleHref = eachRole.Links.Self.HRef
		} else if eachRole.Label.Name == memberRole {
			memberRoleHref = eachRole.Links.Self.HRef
		}
	}

	if adminRoleHref == "" || memberRoleHref == "" {
		return fmt.Errorf("AssignGroupsToProject: Failed to get admin and or member role(s) href")
	}

	args := hcpModels.ExternalUserGroupListResource{
		ExternalUserGroups: []hcpModels.ExternalUserGroupResource{
			{
				Role:  adminRoleHref,
				Group: projectOwnerClaimValue + "-" + projectID,
			},
			{
				Role:  memberRoleHref,
				Group: projectMemberClaimValue + "-" + projectID,
			},
		},
	}

	requestBody, _ := json.Marshal(args)
	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1+epicProjectID+externalUserGroupsFlag,
		http.MethodPut,
		map[string]string{sessionHeader: session},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return errors.Wrapf(err, "AssignGroupsToProject: Failed to Add Groups in HCP: %v", err)
	}

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return err
	}

	if code := resp.StatusCode; code != http.StatusNoContent {
		return fmt.Errorf("AssignGroupsToProject: incorrect status code for assign groups to project %v: %d", epicProjectID, resp.StatusCode)
	}

	return nil
}

// GetConfig is a function which is used to fetch site wide configuration from HCP
func (c *ClientImpl) GetConfig(ctx context.Context, hcpHostURL string) (hcpModels.HcpConfig, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Get Config")
	defer span.Finish()

	session, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return hcpModels.HcpConfig{}, err
	}

	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Fetch Config")
	defer func() { monitor.RecordWithStatus(status) }()

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+configPathV1,
		http.MethodGet,
		map[string]string{sessionHeader: session},
		bytes.NewReader(nil),
	)
	if err != nil {
		return hcpModels.HcpConfig{}, errors.Wrapf(err, "while fetching config from MLOps controller platform.")
	}
	resp.Body.Close()

	status = Success

	err = c.deleteSession(ctx, hcpHostURL, session)
	if err != nil {
		return hcpModels.HcpConfig{}, err
	}

	configResp := hcpModels.HcpConfig{}
	err = json.NewDecoder(resp.Body).Decode(&configResp)
	if err != nil {
		log.Errorf("Error while decoding the value from HCP. Err: %v", err)
		return hcpModels.HcpConfig{}, err
	}

	return configResp, nil
}

//RemoveSiteAdminClaimInHCP updates the site admin project user group
func (c *ClientImpl) RemoveSiteAdminClaimInHCP(ctx context.Context, tenantID, applianceID, hcpHostURL string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Remove Site Admin Claim in Hcp")
	defer span.Finish()
	// Get List of Tenants/projects from HCP. Get Tenant Id for which matches "Site Admin" as tenant Name"
	status = Failure
	monitor := metrics.StartExternalCall(externalSvcName, "Delete SiteAdmin Claim in HCP")
	defer func() { monitor.RecordWithStatus(status) }()

	sessionID, err := c.getSession(ctx, hcpHostURL, hcpUserName, hcpPassword)
	if err != nil {
		return errors.Wrapf(err, "while getting session for delete Site Admin User Group.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}

	resp, err := mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+projectPathV1,
		http.MethodGet,
		map[string]string{sessionHeader: sessionID},
		bytes.NewReader(nil),
	)
	if err != nil {
		return errors.Wrapf(err, "while fetching projects from MLOps controller platform.")
	}

	var projects hcpModels.ListTenants
	_, parseRespErr := common.ParseResponse(resp, &projects)
	if parseRespErr != nil {
		log.Errorf("Failed to fetch projects from HCP: %v", parseRespErr)
		return errors.Wrapf(parseRespErr, "Failed to fetch projects from HCP")
	}

	var siteAdminProjectHRef string
	for _, project := range projects.Embedded.Tenants {
		if project.Label.Name == siteAdminProject {
			siteAdminProjectHRef = project.Links.Self.Href
			break
		}
	}
	if siteAdminProjectHRef == "" {
		return fmt.Errorf("failed while getting site admin project ref from Hcp.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}

	/* Removing UserGroup */
	args := hcpModels.ExternalUserGroupListResource{}
	requestBody, _ := json.Marshal(args)

	resp, err = mlopsHttp.ExecuteHTTPRequest(
		ctx,
		c.client,
		hcpHostURL+siteAdminProjectHRef+externalUserGroupsFlag,
		http.MethodPut,
		map[string]string{sessionHeader: sessionID},
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return errors.Wrapf(err, "Deleting User group in Site Admin Project in HCP failed: %v for TenantID: %s ApplianceID: %s", tenantID, applianceID, err)
	}
	if code := resp.StatusCode; code != http.StatusNoContent {
		return fmt.Errorf("Unable to delete User group for Site Admin project %v: %d.TenantID: %s ApplianceID: %s", siteAdminProjectHRef, resp.StatusCode, tenantID, applianceID)
	}

	status = Success
	err = c.deleteSession(ctx, hcpHostURL, sessionID)
	if err != nil {
		return errors.Wrapf(err, "while deleting session for update Site Admin User Group.TenantID: %s ApplianceID: %s", tenantID, applianceID)
	}
	return nil
}
