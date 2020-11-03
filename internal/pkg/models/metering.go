// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// MeterRequest is struct to have all required attributes
type MeterRequest struct {
	TenantID          string `json:"tenant_id,omitempty"`
	ServiceInstanceID string `json:"service_instance_id,omitempty"`
	LocationID        string `json:"location_id,omitempty"`
	CollectionDate    string `json:"collection_date,omitempty"`
	MeterType         string `json:"resource_type,omitempty"`
	InitiatedBy       string `json:"initiated_by,omitempty"`
	StartDate         string `json:"start_date,omitempty"`
	NumDays           string `json:"num_days,omitempty"`
}

// MeterResponse is struct to have mlops-metering attributes
type MeterResponse struct {
	TenantID          string `json:"tenant_id"`
	ServiceInstanceID string `json:"service_instance_id"`
	ResourceType      string `json:"resource_type"`
	CollectionDate    string `json:"collection_date"`
	LocationID        string `json:"location_id"`
	Status            string `json:"status"`
	Description       string `json:"description"`
	CollectedAt       string `json:"collected_at"`
	InitiatedBy       string `json:"initiated_by"`
}

// MeteringResponse is response struct for the request
type MeteringResponse struct {
	CollectResponse []*MeterResponse
}

// AuditRequest is struct to hold all the filters to fetch audit history
type AuditRequest struct {
	TenantID          string `json:"tenant_id,omitempty"`
	ServiceInstanceID string `json:"service_instance_id,omitempty"`
	ResourceType      string `json:"resource_type,omitempty"`
	StartDate         string `json:"start_date,omitempty"`
	NumDays           int32  `json:"num_days,omitempty"`
}
