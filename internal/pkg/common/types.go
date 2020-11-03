// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

//go:generate mockgen -source ./types.go -package common -destination ./types_mock.go

// Package common - http interface package for this project
package common

import "net/http"

// HTTPClient interface implements HTTP methods
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
