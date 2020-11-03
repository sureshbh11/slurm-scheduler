// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package saml

//Application application object for saml client
type Application struct {
	ID        string
	Kid       string
	Name      string
	SpaceID   string
	SSOAcsURL string
}

//ClaimMappings claimappings for saml client
type ClaimMappings struct {
	ClaimValue  string
	ResourceID  string
	Permissions []string
}

type patchReq struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type createApplicationRes struct {
	createApplicationReq
	ID          string      `json:"id"`
	Credentials credentials `json:"credentials"`
}

type credentials struct {
	Signing struct {
		Kid string `json:"kid"`
	} `json:"signing"`
}

type createApplicationReq struct {
	Settings settings `json:"settings"`
	SpaceID  string   `json:"spaceID"`
	Name     string   `json:"label"`
}

type settings struct {
	SignOn signOn `json:"signOn"`
}

type signOn struct {
	DefaultRelayState     string   `json:"defaultRelayState"`
	SSOAcsURL             string   `json:"ssoAcsUrl"`
	Audience              string   `json:"audience"`
	Recipient             string   `json:"recipient"`
	Destination           string   `json:"destination"`
	SubjectNameIDTemplate string   `json:"subjectNameIdTemplate"`
	SubjectNameIDFormat   string   `json:"subjectNameIdFormat"`
	ResponseSigned        bool     `json:"responseSigned"`
	AssertionSigned       bool     `json:"assertionSigned"`
	SignatureAlgorithm    string   `json:"signatureAlgorithm"`
	DigestAlgorithm       string   `json:"digestAlgorithm"`
	HonorForceAuthn       bool     `json:"honorForceAuthn"`
	AuthContextClassRef   string   `json:"authnContextClassRef"`
	InlineHooks           []string `json:"inlinehooks"`
	AttributeStatements   []string `json:"attributeStatements"`
}

type createAppAssigner struct {
	SpaceID       string     `json:"spaceId"`
	ApplicationID string     `json:"applicationId"`
	Name          string     `json:"name"`
	Permissions   []string   `json:"permissions"`
	Resources     []resource `json:"resources"`
}

type resource struct {
	ID   string `json:"id"`
	ABID string `json:"abId"`
}

type createAppClaimMapping struct {
	SpaceID       string         `json:"spaceId"`
	ApplicationID string         `json:"applicationId"`
	Name          string         `json:"name"`
	ClaimMapping  []claimMapping `json:"claimMapping"`
}

type claimMapping struct {
	Name        string       `json:"name"`
	MatchType   string       `json:"matchType"`
	ClaimValues []claimValue `json:"values"`
}

type claimValue struct {
	Value string `json:"value"`
	When  when   `json:"when"`
}

type when struct {
	AllOf []allOf `json:"allOf"`
}

type allOf struct {
	Match match `json:"match"`
}

type match struct {
	Resource   resource `json:"resource"`
	Permission string   `json:"permission"`
}

type responseID struct {
	ID string `json:"id"`
}
