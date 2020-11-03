// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

import "errors"

// Define errors that callers of the DataStore should use
// TODO: have a second pass at it to turn these into structs with detailsed error codes
var (
	ErrServiceInstanceNotFound = errors.New("The specified service instance was not found in the dynamodb")
	ErrProjectNotFound         = errors.New("The specified project was not found in the dynamodb")
)

// AWSCreds holds configuration values for DB Connection
type AWSCreds struct {
	DBAwsAccessKeyID     string `env:"DB_AWS_ACCESS_KEY_ID"`
	DBAwsSecretAccessKey string `env:"DB_AWS_SECRET_ACCESS_KEY"`
	DBRegionName         string `env:"DB_AWS_REGION_NAME" envDefault:"us-west-2"`
	DBEndpointURL        string `env:"DB_AWS_ENDPOINT" envDefault:"dynamodb.amazonaws.com"`
	DBTableName          string `env:"DB_AWS_TABLE_NAME" envDefault:"mlops-osb-dev"`
}

// ServiceInstance is a struct contains DB Instance Information
type ServiceInstance struct {
	ID                  string                 `json:"id"`
	Type                string                 `json:"type"`
	TenantID            string                 `json:"tenantID,omitempty"`
	ServiceID           string                 `json:"serviceID"`
	PlanID              string                 `json:"planID"`
	Status              string                 `json:"status,omitempty"`
	TrialID             string                 `json:"trialID,omitempty"`
	OrganizationGUID    string                 `json:"organizationGUID,omitempty"`
	SpaceGUID           string                 `json:"spaceGUID,omitempty"`
	Parameters          map[string]interface{} `json:"parameters,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Output              string                 `json:"output,omitempty"`
	LocationID          string                 `json:"locationID,omitempty"`
	LocationName        string                 `json:"locationName,omitempty"`
	LocationDescription string                 `json:"locationDescription,omitempty"`
	StartDate           string                 `json:"startDate,omitempty"`
	EndDate             string                 `json:"endDate,omitempty"`
	CreatedAt           string                 `json:"createdAt,omitempty"`
	UpdatedAt           string                 `json:"updatedAt,omitempty"`
	DeletedAt           string                 `json:"deletedAt,omitempty"`
	Version             string                 `json:"version,omitempty"`
}

// ServiceBinding - Holds Service binding details
type ServiceBinding struct {
	ID                string              `json:"id"`
	Type              string              `json:"type"`
	TenantID          string              `json:"tenantID,omitempty"`
	PlanID            string              `json:"planID"`
	BindingURLs       []ServiceBindingURL `json:"bindingURLs"`
	ServiceID         string              `json:"serviceID"`
	ServiceInstanceID string              `json:"serviceInstanceID"`
	CreatedAt         string              `json:"createdAt,omitempty"`
	UpdatedAt         string              `json:"updatedAt,omitempty"`
	DeletedAt         string              `json:"deletedAt,omitempty"`
	Version           string              `json:"version,omitempty"`
	Kubeconfig        string              `json:"kubeconfig,omitempty"`
}

// ServiceBindingURL - Holds binding urls
type ServiceBindingURL struct {
	Service    string `json:"service"`
	BindingURL string `json:"bindingURL"`
	TypeURL    string `json:"typeURL"`
}
