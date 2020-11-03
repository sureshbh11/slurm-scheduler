// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

//CommonLabel contains the name a description of the object which contains it.
type CommonLabel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//CommonLink contains a link to the object which contains it.
type CommonLink struct {
	HRef string `json:"href"`
}

// CommonResourceLinks contains a link to the containing resource path.
type CommonResourceLinks struct {
	Self CommonLink `json:"self"`
}

// RoleListEmbeddedResource contains information for multiple roles.
type RoleListEmbeddedResource struct {
	Roles []RoleResource `json:"roles"`
}

// RoleListResource is the return value for the Get API call.
type RoleListResource struct {
	Links    CommonResourceLinks      `json:"_links"`
	Embedded RoleListEmbeddedResource `json:"_embedded"`
}

// RoleResource contains a role resource definition.
type RoleResource struct {
	Label CommonLabel         `json:"label"`
	Links CommonResourceLinks `json:"_links"`
}
