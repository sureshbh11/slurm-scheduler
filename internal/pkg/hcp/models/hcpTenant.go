// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// ListTenants model which specifies list of tenants
type ListTenants struct {
	Links    Links    `json:"_links"`
	Embedded Embedded `json:"_embedded"`
}

// Embedded model which specifies upper level struct which consists of list of tenants
type Embedded struct {
	Tenants []Tenant `json:"tenants"`
}

// Tenant model which specifies HCP tenant
type Tenant struct {
	Links                       Links               `json:"_links"`
	Label                       Label               `json:"label"`
	MemberKeyAvailable          string              `json:"member_key_available"`
	ConstraintsSupported        bool                `json:"constraints_supported"`
	ClusterIsolationSupported   bool                `json:"cluster_isolation_supported"`
	GPUUsageSupported           bool                `json:"gpu_usage_supported"`
	FilesystemMountSupported    bool                `json:"filesystem_mount_supported"`
	PersistentSupported         bool                `json:"persistent_supported"`
	Status                      string              `json:"status"`
	Features                    Features            `json:"features"`
	Quota                       Quota               `json:"quota"`
	Inusequota                  Quota               `json:"inusequota"`
	TenantStorageQuotaSupported bool                `json:"tenant_storage_quota_supported"`
	ExternalUserGroups          []ExternalUserGroup `json:"external_user_groups"`
	TenantEnforcements          []interface{}       `json:"tenant_enforcements"`
	QosMultiplier               *int64              `json:"qos_multiplier,omitempty"`
	TenantType                  *string             `json:"tenant_type,omitempty"`
	TenantConstraints           []TenantConstraint  `json:"tenant_constraints"`
	TenantTags                  []TenantTag         `json:"tenant_tags"`
}

// ExternalUserGroup model which specifies external user groups related to tenant
type ExternalUserGroup struct {
	Group string `json:"group"`
	Role  string `json:"role"`
}

// Features model which specifies features related to tenant
type Features struct {
	MlProject        bool `json:"ml_project"`
	KubernetesAccess bool `json:"kubernetes_access"`
}

// Quota model which specifies quotas related to tenant
type Quota struct {
	Cores         *int64 `json:"cores,omitempty"`
	Gpus          *int64 `json:"gpus,omitempty"`
	Memory        *int64 `json:"memory,omitempty"`
	TenantStorage *int64 `json:"tenant_storage,omitempty"`
	Disk          *int64 `json:"disk,omitempty"`
	Persistent    *int64 `json:"persistent,omitempty"`
}

// Label model which specifies Laeblling related to tenant
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Links specifies the uri's  that links to various other resources
type Links struct {
	Self Self `json:"self"`
}

// Self model which specifies its own uri link
type Self struct {
	Href string `json:"href"`
}

// TenantConstraint model which specifies constraint related to tenant
type TenantConstraint struct {
	Operator string `json:"operator"`
	Tag      string `json:"tag"`
	TagValue string `json:"tag_value"`
}

// TenantTag model which specifies tags related to tenant
type TenantTag struct {
	Tag           string   `json:"tag"`
	AllowedValues []string `json:"allowed_values"`
}
