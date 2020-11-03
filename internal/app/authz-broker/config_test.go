// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package authzbroker

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testExternalURL     = "http://testexternal.com"
	testScmConfigPath   = "test_config/path"
	testIamURL          = "https://test.iam.com"
	testHPCaasABPort     = "4040"
	expectedHPCaasABPort = 4040
)

func TestNewConfigSuccess(t *testing.T) {

	os.Setenv("EXTERNALURL", testExternalURL)
	os.Setenv("SCM_CONFIG_PATH", testScmConfigPath)
	os.Setenv("IAMURL", testIamURL)
	os.Setenv("HPCaas_AB_PORT", testHPCaasABPort)

	cfg, err := NewConfig()

	assert.Nil(t, err, "nil error is expected")
	assert.Equal(t, testExternalURL, cfg.ExternalURL)
	assert.Equal(t, testScmConfigPath, cfg.SCMConfigPath)
	assert.Equal(t, testIamURL, cfg.IamURL)
	assert.Equal(t, expectedHPCaasABPort, cfg.Port)
}

func TestNewConfigError(t *testing.T) {

	os.Setenv("EXTERNALURL", testExternalURL)
	os.Setenv("SCM_CONFIG_PATH", testScmConfigPath)
	os.Setenv("IAMURL", testIamURL)
	os.Setenv("HPCaas_AB_PORT", "error")

	_, err := NewConfig()

	assert.NotNil(t, err, "error is expected")
}
