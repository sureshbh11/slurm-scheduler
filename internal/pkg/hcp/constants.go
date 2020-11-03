// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package hcp

const (
	sessionHeader           = "X-BDS-SESSION"
	configPath              = "/api/v2/config/auth"
	sessionPath             = "/api/v2/session/"
	projectPathV1           = "/api/v1/tenant/"
	configPathV1            = "/api/v1/config"
	projectPathV2           = "/api/v2/tenant/"
	clusterPathV2           = "/api/v2/cluster/"
	modelPathV2             = "/api/v2/model/"
	rolePathV1              = "/api/v1/role"
	adminRole               = "Admin"
	memberRole              = "Member"
	siteAdminRole           = "Site Admin"
	siteAdminProject        = "Site Admin"
	externalUserGroupsFlag  = "?external_user_groups"
	projectOwnerClaimValue  = "saml-project-admin"
	projectMemberClaimValue = "saml-project-member"
	siteAdminClaimValue     = "saml-site-admin"
)
