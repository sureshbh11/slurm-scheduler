// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package models

//Header struct contains saml ID details
type Header struct {
	ID             string `json:"id"`   // uuid generated during create
	Name           string `json:"name"` // names are unique per tenant
	CreatedDate    string `json:"createdDate"`
	LastUpdateDate string `json:"lastUpdateDate"`
}

// ApplianceInput struct contains input required for SAML configuration
type ApplianceInput struct {
	Name      string `json:"name"`
	SSOAcsURL string `json:"ssoAcsUrl"`
}

// HcpSAMLInput struct contains input required to configure HCP SAML
type HcpSAMLInput struct {
	TenantID  string         `json:"tenantID"`
	Appliance ApplianceInput `json:"appliance"`
}

//Appliance rest api model
type Appliance struct {
	Header
	Saml          *SAMLMetaData `json:"samlMetaData"`
	TenantID      string        `json:"tenantID"`
	SiteID        string        `json:"siteID"`
	ApplicationID string        `json:"-"`
	AppAssignerID string        `json:"-"`
	IamAPIVersion string        `json:"-"`
}

//SAMLMetaData  saml metadata
type SAMLMetaData struct {
	SSOAcsURL       string `json:"ssoAcsUrl"`
	Kid             string `json:"-"`
	ClaimMappingsID string `json:"-"`
}

// ExternalIdentityServer show type of Server
type ExternalIdentityServer struct {
	Type             string `json:"type"`
	VerifyPeer       bool   `json:"verify_peer,omitempty"`
	BindType         string `json:"bind_type,omitempty"`
	Host             string `json:"host,omitempty"`
	Port             int    `json:"port,omitempty"`
	BaseDN           string `json:"base_dn,omitempty"`
	BindDN           string `json:"bind_dn,omitempty"`
	BindPwd          string `json:"bind_pwd,omitempty"`
	UserAttribute    string `json:"user_attribute,omitempty"`
	SecurityProtocol string `json:"security_protocol,omitempty"`
}

// SsoSettings contains sso configuations related to SAML App
type SsoSettings struct {
	PruneSubject            bool   `json:"prune_subject"`
	SamlUserXpath           string `json:"saml_user_xpath"`
	AllowUnPwLogin          bool   `json:"allow_un_pw_login"`
	PreventExtUserUnPwLogin bool   `json:"prevent_ext_user_un_pw_login"`
	SamlMetadata            string `json:"saml_metadata"`
	SamlMetadataFilename    string `json:"saml_metadata_filename"`
	SamlGroupXpath          string `json:"saml_group_xpath"`
}

// EpicConfigInput contains parameters required configure EPIC Controller
type EpicConfigInput struct {
	ExternalServer  ExternalIdentityServer `json:"external_identity_server"`
	EpicSsoSettings SsoSettings            `json:"sso_settings"`
}

// MetadataOutput contains metadata config for the given SAML App
type MetadataOutput struct {
	SamlAppID string `json:"appID"`
	Metadata  string `json:"metadata"`
}
