//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"github.com/gin-gonic/gin"

	hpeErrors "github.com/hpe-hcss/errors/pkg/errors"
	"github.com/hpe-hcss/iam-lib/pkg/resource/authorization"
	"github.com/hpe-hcss/loglib/pkg/log"
)

// AuthorizationBrokerID is the public broker ID for IAM broker
const AuthorizationBrokerID = "HPCaas"

// ProjectResourceIdentifier generate an IAM specific resource
// identifier for projects in HCP
func ProjectResourceIdentifier(id string) authorization.ResourceIdentifier {
	return authorization.ResourceIdentifier{
		ID:                    "projects/" + id,
		AuthorizationBrokerID: AuthorizationBrokerID,
	}
}

// ApplianceResourceIdentifier generate an IAM specific resource
// identifier for appliances in HCP
func ApplianceResourceIdentifier(id string) authorization.ResourceIdentifier {
	return authorization.ResourceIdentifier{
		ID:                    "appliances/" + id,
		AuthorizationBrokerID: AuthorizationBrokerID,
	}
}

// ProcessAuthorizeError processes the error from iam.Authorize and set's the
// response on the context. This is a helper function used in places
// where iam.Authorize is called.
func ProcessAuthorizeError(c *gin.Context, err error, tenant, userID, permission string, resource string) {
	_, forbidden := err.(*hpeErrors.ErrForbidden)
	if forbidden {
		log.Infof(c.Request.Context(), "Tenant %v User %v does not have permission %v on resource %v", tenant, userID, permission, resource)
	} else {
		log.Errorf(c.Request.Context(), "Error in permission check tenantID %v userID %v permission %v on resource %v: %v",
			tenant, userID, permission, resource, err)
	}
	hpeErrors.SetResponseIfError(c, err)
}
