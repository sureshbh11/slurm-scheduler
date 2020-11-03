//(c) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"context"

	"github.com/hpe-hcss/iam-lib/pkg/token"
)

// GetAccessToken retrieve token from request context
func GetAccessToken(ctx context.Context) string {
	return Caller(ctx).Token
}

// Caller get the details of the client that called the API
func Caller(ctx context.Context) token.AuthTokenDetails {
	return token.GetAuthTokenDetailsFromContext(ctx)
}
