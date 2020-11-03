//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// DeleteProjectInput is the payload required to delete project from hcp
type DeleteProjectInput struct {
	TenantID          string `json:"tenant_id"`
	UserID            string `json:"user_id"`
	ProjectID         string `json:"project_id"`
	ServiceInstanceID string `json:"serviceInstanceID"`
}

// GetProjectInput is the payload required to get project from hcp
type GetProjectInput struct {
	TenantID          string `json:"tenant_id" validate:"IDValidator"`
	UserID            string `json:"user_id" validate:"IDValidator"`
	ProjectID         string `json:"project_id"`
	ServiceInstanceID string `json:"serviceInstanceID"`
	SpaceID           string `json:"spaceID" validate:"IDValidator"`
}

// CreateProjectInput structure which defines the parameters required for on creating project in hcp
type CreateProjectInput struct {
	Quota             Quota  `json:"quota"`
	Label             Label  `json:"label"`
	ServiceInstanceID string `json:"serviceInstanceID"`
	SpaceName         string `json:"spaceName"`
	SpaceID           string `json:"spaceID" validate:"IDValidator"`
	TenantID          string `json:"tenant_id" validate:"IDValidator"`
	UserID            string `json:"user_id" validate:"IDValidator"`
}

// Project is structure which represents top level Project Structure
type Project struct {
	Header
	CreateProjectInput
	EPICTenantID string `json:"epicTenantID"`
	SiteName     string `json:"siteName"`
}

// Quota structure which defines the quota of a Project
type Quota struct {
	Cores             int64 `json:"cores"`
	Gpus              int64 `json:"gpus,omitempty"`
	TenantStorage     int64 `json:"tenant_storage,omitempty"`
	Memory            int64 `json:"memory,omitempty"`
	PersistentStorage int64 `json:"persistent,omitempty"`
	Disk              int64 `json:"disk,omitempty"`
}

// HCPQuota structure which defines the quota of a HCP project
type HCPQuota struct {
	Cores             *int64 `json:"cores"`
	Gpus              *int64 `json:"gpus,omitempty"`
	TenantStorage     *int64 `json:"tenant_storage,omitempty"`
	Memory            *int64 `json:"memory,omitempty"`
	PersistentStorage *int64 `json:"persistent,omitempty"`
	Disk              *int64 `json:"disk,omitempty"`
}

// Label structure defines the naming parameters
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Features structure defines the extra features for a project
type Features struct {
	MlProject        bool `json:"ml_project"`
	KubernetesAccess bool `json:"kubernetes_access"`
}

// Resources structure defines the quota model supported by a project
type Resources struct {
	Cores                   int64 `json:"cores"`
	Gpus                    int64 `json:"gpus"`
	MemoryGb                int64 `json:"memory_gb"`
	DiskGb                  int64 `json:"disk_gb"`
	TenantStorageGb         int64 `json:"tenant_storage_gb"`
	PersistentStorageSizeGb int64 `json:"persistent_storage_size_gb"`
}

// TenantEnclosedProperties structure defines the properties of a project
type TenantEnclosedProperties struct {
	QosMultiplier int `json:"qos_multiplier"`
}

// ProjectOutput structure defines the project model from hcp
type ProjectOutput struct {
	ID         string    `json:"id"`
	Label      Label     `json:"label"`
	Quota      Resources `json:"quota,omitempty"`
	InUseQuota Resources `json:"inusequota,omitempty"`
	Site       string    `json:"siteName"`
	TenantID   string    `json:"tenantID"`
}

// HCPCreateProjectInput structure which defines the parameters required for on creating project in hcp
type HCPCreateProjectInput struct {
	QosMultiplier      int      `json:"qos_multiplier"`
	MemberKeyAvailable string   `json:"member_key_available"`
	TenantType         string   `json:"tenant_type"`
	Quota              HCPQuota `json:"quota,omitempty"`
	Label              Label    `json:"label"`
	Features           Features `json:"features"`
}

// GetProjectKPIOutput structure defines the KPI model from hcp
type GetProjectKPIOutput struct {
	NumberOfTrainingClusters   int32 `json:"numberOfTrainingClusters"`
	NumberOfDeploymentClusters int32 `json:"numberOfDeploymentClusters"`
	NumberOfNotebookServers    int32 `json:"numberOfNotebookServers"`
	NumberOfModelRegistry      int32 `json:"numberOfModelRegistry"`
}

// GetSiteKPIOutput structure defines the KPI model from hcp
type GetSiteKPIOutput struct {
	NumberOfProjects int32 `json:"numberOfProjects"`
}

// HCPResources steucture defines the list of HCP level resources available
type HCPResources struct {
	TotalCores             int64   `json:"totalCores"`
	TotalGpus              int64   `json:"totalGpus"`
	TotalMemory            float64 `json:"totalMemory"`
	TotalDisk              float64 `json:"totalDisk"`
	TotalTenantStorage     float64 `json:"totalTenantStorage"`
	TotalPersistentStorage float64 `json:"totalPersistentStorage"`
}
