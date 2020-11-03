//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package abregistrar

import (
	"github.com/hpe-hcss/iam-lib/pkg/resource"
)

const (
	alphaVersion = "alpha"
	//HPCaasClusterManage permission to manage cluster
	HPCaasClusterManage = "HPCaas.cluster.manage"
)

var (
	permissions = map[string]resource.Permission{
		"HPCaas.server.create": {
			ID:          "HPCaas.server.create",
			Name:        "Create HPCaas Server",
			Description: "Create HPCaas Server",
			Version:     alphaVersion,
		},
		"HPCaas.server.read": {
			ID:          "HPCaas.server.read",
			Name:        "Read HPCaas Server",
			Description: "Read HPCaas Server",
			Version:     alphaVersion,
		},
		"HPCaas.server.update": {
			ID:          "HPCaas.server.update",
			Name:        "Update HPCaas Server",
			Description: "Update HPCaas Server",
			Version:     alphaVersion,
		},
		"HPCaas.server.delete": {
			ID:          "HPCaas.server.delete",
			Name:        "Delete HPCaas Server",
			Description: "Delete HPCaas Server",
			Version:     alphaVersion,
		},
		"HPCaas.unilogin.datacenteradmin": {
			ID:          "HPCaas.unilogin.datacenteradmin",
			Name:        "HPCaas UniLogin Datacenteradmin",
			Description: "HPCaas UniLogin Datacenteradmin",
			Version:     alphaVersion,
		},
		"HPCaas.unilogin.serveradmin": {
			ID:          "HPCaas.unilogin.serveradmin",
			Name:        "HPCaas UniLogin Serveradmin",
			Description: "HPCaas UniLogin Serveradmin",
			Version:     alphaVersion,
		},
		"HPCaas.unilogin.servermember": {
			ID:          "HPCaas.unilogin.servermember",
			Name:        "HPCaas UniLogin Servermember",
			Description: "HPCaas UniLogin Servermember",
			Version:     alphaVersion,
		},
	}
)
