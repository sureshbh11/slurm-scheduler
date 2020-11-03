//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package abregistrar

import (
	"github.com/hpe-hcss/iam-lib/pkg/resource"
)

var roles = map[string]resource.Role{
	"HPCaas.datacenter.admin": {
		Name:   "HPCaas Datacenter Admin",
		Source: "PREDEFINED",
		Permissions: []string{
			"HPCaas.unilogin.datacenteradmin",
		},
	},
	"HPCaas.admin": {
		Name:   "HPCaas Admin",
		Source: "PREDEFINED",
		Permissions: []string{
			"HPCaas.server.read",
			"HPCaas.server.create",
			"HPCaas.server.update",
			"HPCaas.server.delete",
		},
	},
	"HPCaas.server.admin": {
		Name:   "HPCaas Server Admin",
		Source: "PREDEFINED",
		Permissions: []string{
			"HPCaas.server.read",
			"HPCaas.unilogin.serveradmin",
		},
	},
	"HPCaas.server.member": {
		Name:   "HPCaas Server Member",
		Source: "PREDEFINED",
		Permissions: []string{
			"HPCaas.server.read",
			"HPCaas.unilogin.servermember",
		},
	},
}
