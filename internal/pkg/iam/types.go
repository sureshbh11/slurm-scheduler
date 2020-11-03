//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"context"

	"github.com/hpe-hcss/iam-lib/pkg/resource/authorization"
)

const (
	// HPCaasProjectRead an IAM permission for reading projects
	HPCaasProjectRead = "HPCaas.project.read"

	// HPCaasProjectDelete an IAM permission for deleting project
	HPCaasProjectDelete = "HPCaas.project.delete"

	// HPCaasProjectUpdate an IAM permission for updating project
	HPCaasProjectUpdate = "HPCaas.project.update"

	// HPCaasProjectCreate an IAM permission to create a project
	HPCaasProjectCreate = "HPCaas.project.create"

	// HPCaasSiteRead an IAM permission for reading job
	HPCaasSiteRead = "HPCaas.site.read"
)

// Registrar registers permissions with the IAM system
type Registrar interface {
	RegisterAuthorizationBroker(ctx context.Context,
		tenantID string,
		token string,
		spaceID string,
		externalURL string) error

	AddRoles(ctx context.Context,
		tenantID string,
		token string,
		spaceID string) error

	DeRegisterAuthorizationBroker(ctx context.Context,
		tenantID string,
		token string,
		spaceID string) error

	DeleteRoles(ctx context.Context,
		tenantID string,
		token string) error
}

// Authorizer is an interface to the IAM system
type Authorizer interface {
	Authorize(ctx context.Context, tenant, permission, userID string,
		resource authorization.ResourceIdentifier) error
}
