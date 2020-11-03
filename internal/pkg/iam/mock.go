//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package iam

import (
	"context"
	"github.com/hpe-hcss/iam-lib/pkg/resource/authorization"
)

// Mock is a mock implementation of the IAM interface
type Mock struct{}

// Authorize returns true
func (i Mock) Authorize(ctx context.Context, accessToken, permission, userID string,
	resource authorization.ResourceIdentifier) error {
	return nil
}

// RegisterPermissions registers permissions with the IAM system
func (i Mock) RegisterPermissions() error {
	return nil
}

// NewMock returns an IAM Authorizer
func NewMock() (Authorizer, error) {
	return Mock{}, nil
}
