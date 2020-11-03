// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

//go:generate mockgen -source ./client.go -package authzbroker -destination ./client_mock.go

// Package authzbroker - main package for authzbroker
package authzbroker

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	apiErrors "github.com/hpe-hcss/errors/pkg/errors"
	abres "github.com/hpe-hcss/iam-lib/pkg/resource"
	spacesclient "github.com/hpe-hcss/iam-lib/pkg/spaces-client"
	"github.com/hpe-hcss/loglib/pkg/errors"
	"github.com/hpe-hcss/loglib/pkg/log"
	"github.com/hpe-hcss/scm-lib/pkg/scmclient"

	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/constants"
	abregistrar "github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/iam/ab-registrar"
	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/utils"
)

const (
	tenantIDAuthFailed = "TenantID authorization failed. %v"
	tenantServicePerm  = "Need tenant.service.manager permission on %s"
	degisterAbFailed   = "unable to de-register AB with tenant %v"
)

type serviceAuthClient interface {
	IsAuthorizedToProvisionServicesForTenant(context.Context, string, string) error
}

// AuthBrokerGRPC interface contains GRPC methods required for AuthBroker Operations
type AuthBrokerGRPC interface {
	RegisterAuthBroker(context.Context, string) error
	DeRegisterAuthBroker(context.Context, string) error
}

// AuthorizationBroker type wraps the authorization broker functions
type AuthorizationBroker struct {
	serviceConfig ServiceOptions
	scmClient     scmclient.SCMAPI
	abClient      abregistrar.Registrar
	spacesClient  spacesclient.IamSpacesAPI
	spaceName     string
}

type tenant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Broker returns a new authorization broker object
func Broker(cfg ServiceOptions,
	scmClient scmclient.SCMAPI,
	abClient abregistrar.Registrar,
	spacesClient spacesclient.IamSpacesAPI,
	spaceName string,
) *AuthorizationBroker {

	return &AuthorizationBroker{
		serviceConfig: cfg,
		scmClient:     scmClient,
		abClient:      abClient,
		spacesClient:  spacesClient,
		spaceName:     spaceName,
	}
}

// BootStrapAB looks for new tenants and registers the AuthorizationBroker for each new tenant.
func (a *AuthorizationBroker) BootStrapAB(c *gin.Context) {
	ctx := c.Request.Context()

	tenantReqBody := tenant{}
	err := c.BindJSON(&tenantReqBody)
	if err != nil {
		log.Error(ctx, err)
		c.Status(http.StatusBadRequest)
		return
	}

	log.Infof(ctx, "register Authorization broker with %v", tenantReqBody.ID)

	err = a.RegisterAuthBroker(ctx, tenantReqBody.ID)
	if err != nil {
		log.Error(ctx, err)
		apiErrors.SetResponseIfError(c, err)
		return
	}
	c.Status(http.StatusOK)
}

// Resources returns all HPCaas resources
func (a *AuthorizationBroker) Resources(c *gin.Context) {
	resources := abres.BrokerResponse{
		Path: abres.Path{
			Name: constants.ContainerPlatform,
			Path: constants.ContainerPlatformPath,
			Leaf: false,
		},
		SubPaths: []abres.Path{
			{
				Name: constants.RootAppliance,
				Path: constants.ContainerPlatformPath + constants.RootAppliancePath,
				Leaf: false,
			},
			{
				Name: constants.RootProject,
				Path: constants.ContainerPlatformPath + constants.RootProjectPath,
				Leaf: false,
			},
		},
	}
	c.JSON(http.StatusOK, resources)
}

// RegisterAuthBroker method registers AuthBroker including HPCaas Roles for the given tenant
func (a *AuthorizationBroker) RegisterAuthBroker(ctx context.Context, tenantID string) error {

	token, err := a.scmClient.VendTenantToken(ctx, tenantID)
	if err != nil {
		return errors.Wrapf(err, "unable to vend a token for tenant '%v'", tenantID)
	}

	spaceID, err := utils.GetSpaceID(ctx, a.spacesClient, token, a.spaceName)
	if err != nil {
		return errors.Wrapf(err, "unable get spaceID for %v", a.spaceName)
	}

	err = a.abClient.RegisterAuthorizationBroker(ctx, tenantID, token, spaceID, a.serviceConfig.ExternalURL)
	if err != nil {
		return errors.Wrapf(err, "unable to register AB with tenant %v", tenantID)
	}

	err = a.abClient.AddRoles(ctx, tenantID, token, spaceID)
	if err != nil {
		return errors.Wrapf(err, "unable to register roles with tenant %v", tenantID)
	}

	return nil
}

// DeRegisterAB is amethod to delete all roles/permissions/resources and AB from a tenant
func (a *AuthorizationBroker) DeRegisterAB(c *gin.Context) {

	ctx := c.Request.Context()
	tenantReqBody := tenant{}
	err := c.BindJSON(&tenantReqBody)
	if err != nil {
		log.Error(ctx, err)
		c.Status(http.StatusBadRequest)
		return
	}

	err = a.DeRegisterAuthBroker(ctx, tenantReqBody.ID)
	if err != nil {
		log.Error(ctx, err)
		apiErrors.SetResponseIfError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// DeRegisterAuthBroker method de-registers AuthBroker including HPCaas Roles for the given tenant
func (a *AuthorizationBroker) DeRegisterAuthBroker(ctx context.Context, tenantID string) error {

	token, err := a.scmClient.VendTenantToken(ctx, tenantID)
	if err != nil {
		TokenErr := errors.Wrapf(err, "unable to vend a token for tenant '%v'", tenantID)
		log.Error(ctx, TokenErr.Error())
		return TokenErr
	}

	err = a.abClient.DeleteRoles(ctx, tenantID, token)
	if err != nil {
		DelErr := errors.Wrapf(err, "Failed to delete HPCaas roles '%v'", tenantID)
		log.Error(ctx, DelErr.Error())
		return DelErr
	}

	spaceID, err := utils.GetSpaceID(ctx, a.spacesClient, token, a.spaceName)
	if err != nil {
		spaceErr := errors.Wrapf(err, "Unable to get space '%v'", a.spaceName)
		log.Error(ctx, spaceErr.Error())
		return spaceErr
	}

	err = a.abClient.DeRegisterAuthorizationBroker(ctx, tenantID, token, spaceID)
	if err != nil {
		log.Errorf(ctx, degisterAbFailed, err)
		return err
	}
	return nil

}
