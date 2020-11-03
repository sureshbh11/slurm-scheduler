//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// ModelResp structure defines the model from hcp
type ModelResp struct {
	Links    interface{} `json:"_links"`
	Embedded ModelsArray `json:"_embedded"`
}

// ModelsArray structure defines the models array from hcp
type ModelsArray struct {
	Model []Models `json:"models"`
}

// Models structure defines the models from hcp
type Models struct {
	Links Links `json:"_links"`
}

// Links structure defines the links from hcp
type Links struct {
	Tenant Tenant `json:"tenant"`
}

// Tenant structure defines the tenants from hcp
type Tenant struct {
	Href     string   `json:"href"`
	Title    string   `json:"title"`
	Status   string   `json:"status"`
	Features Features `json:"features"`
}
