// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// HcpConfig defines the top level of config from HCP
type HcpConfig struct {
	Result *string `json:"result,omitempty"`
	Config *Config `json:"objects,omitempty"`
}

// Config defines the list of configurations from a particular HCP site
type Config struct {
	TotalTenantStorage *int64        `json:"total_tenant_storage,omitempty"`
	SystemQuota        []SystemQuota `json:"system_quota"`
}

// SystemQuota defines the resources available at HCP Site level
type SystemQuota struct {
	TenantType *string `json:"tenant_type,omitempty"`
	Cores      *int64  `json:"cores,omitempty"`
	Memory     *int64  `json:"memory,omitempty"`
	Swap       *int64  `json:"swap,omitempty"`
	Disk       *int64  `json:"disk,omitempty"`
	Nodes      *int64  `json:"nodes,omitempty"`
	Gpus       *int64  `json:"gpus,omitempty"`
}
