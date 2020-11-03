// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package authzbroker

import (
	"context"

	configEnv "github.com/caarlos0/env"
	"github.com/hpe-hcss/loglib/pkg/log"
)

// ServiceOptions is the set of AZAM AB Service configuration options.
type ServiceOptions struct {
	ExternalURL     string `env:"EXTERNALURL" envDefault:"http://localhost:8080"`
	SCMConfigPath   string `env:"SCM_CONFIG_PATH" envDefault:"scmClientConfig.yaml"`
	IamURL          string `env:"IAMURL" envDefault:"http://localhost:8080"`
	IamGrpc         string `env:"IAMGRPC" envDefault:"http://localhost:443"`
	SCMAWSRegion    string `env:"SCM_AWS_REGION" envDefault:"us-west-2"`
	SCMAWSAccessID  string `env:"SCM_AWS_ACCESS_ID"`
	SCMAWSSecretKey string `env:"SCM_AWS_SECRET_KEY"`
	Port            int    `env:"HPCaas_AB_PORT" envDefault:"9001"`
	GRPCPort        string `env:"HPCaas_AB_GRPC_PORT" envDefault:"5557"`
	AuthType        string `env:"AUTH_TYPE" envDefault:"IAM"`
}

// NewConfig is the factory method that reads and sets default for the HPCaas AB Service.
func NewConfig() (ServiceOptions, error) {
	serviceConfig := ServiceOptions{}

	err := configEnv.Parse(&serviceConfig)
	if err != nil {
		log.Errorf(context.Background(), "Could not parse the service configuration: %v", err)
		return ServiceOptions{}, err
	}

	return serviceConfig, nil
}
