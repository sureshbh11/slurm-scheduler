// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

//go:generate mockgen -source ./interface.go -package abregistrar -destination ./mock_abregistrar.go

package abregistrar

import (
	"context"

	abClient "github.com/hpe-hcss/iam-lib/pkg/ab-registration-client"
	"github.com/hpe-hcss/iam-lib/pkg/resource"
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

// RegistrarClient defines interface to register scheduler with iam
type RegistrarClient interface {
	Register(context.Context, abClient.ClientDetails, resource.AuthorizationBroker, string) error
	AddRole(context.Context, abClient.ClientDetails, string, resource.Role, string) error
	DeleteRole(context.Context, abClient.ClientDetails, string) error
}
