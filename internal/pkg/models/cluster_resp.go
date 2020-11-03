//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// ClusterResp structure defines the clusters from hcp
type ClusterResp struct {
	Links    interface{}   `json:"_links"`
	Embedded ClustersArray `json:"_embedded"`
}

// ClustersArray structure defines the Clusters array from hcp
type ClustersArray struct {
	Clusters []Clusters `json:"clusters"`
}

// Clusters structure defines the clusters from hcp
type Clusters struct {
	Nodegroup Nodegroup `json:"nodegroup"`
	TenantID  string    `json:"tenant_id"`
}

// Nodegroup structure defines the Nodegroup from hcp
type Nodegroup struct {
	IntendedCategories []string `json:"intended_categories"`
}
