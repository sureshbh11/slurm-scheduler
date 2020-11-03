// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

// SessionInput defines the variables required for creating session
type SessionInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// ExternalUserGroupResource contains info about external user group
type ExternalUserGroupResource struct {
	Role  string `json:"role"`
	Group string `json:"group"`
}

// ExternalUserGroupListResource contains info about multiple external user group
type ExternalUserGroupListResource struct {
	ExternalUserGroups []ExternalUserGroupResource `json:"external_user_groups"`
}
