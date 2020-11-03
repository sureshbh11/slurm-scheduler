// Copyright 2020 Hewlett Packard Enterprise Development LP

package utils

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	spacesclient "github.com/hpe-hcss/iam-lib/pkg/spaces-client"
)

var (
	idRegexp = regexp.MustCompile(`^[a-zA-Z0-9][-_a-zA-Z0-9]*[a-zA-Z0-9]$|^[a-zA-Z0-9]$`)
)

// IsValidID is intended to validate parameters and header as part of payload,
// primarily to ensure that you can't pass in /rest/users/../../something
// via double encoding and hit internal URIs you're not meant to.
// An id is valid if it only containsString [-_a-zA-Z0-9]
func IsValidID(id string) bool {
	return idRegexp.MatchString(id)
}

// GetEnv - get the value of an env-var with a fallback
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

//GetSpaceID gets id of given space
func GetSpaceID(ctx context.Context, spacesClient spacesclient.IamSpacesAPI, token, spaceName string) (string, error) {
	spaceList, err := spacesClient.List(ctx, token, spacesclient.ListSpaceInput{Name: spaceName})
	if err != nil {
		return "", errors.Wrapf(err, "unable to get space: %s", spaceName)
	}

	for _, space := range spaceList.Members {
		if space.Name == spaceName {
			return space.ID, nil
		}
	}

	return "", errors.WithStack(fmt.Errorf("no space found for space name %v", spaceName))
}

// GetTenantID gets the project id from tenant_id
func GetTenantID(id string) string {
	stringSplit := strings.Split(id, "/")
	return stringSplit[len(stringSplit)-1]
}
