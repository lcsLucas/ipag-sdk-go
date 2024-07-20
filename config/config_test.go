package config_test

import (
	"testing"

	"github.com/lcslucas/ipag-sdk-go/config"
)

func TestConfig(t *testing.T) {
	expected := struct {
		apiID       string
		apiKey      string
		environment string
		version     uint8
	}{
		apiID:       "123456",
		apiKey:      "ABC123",
		environment: config.Environments.Production,
		version:     2,
	}

	cfg := config.Config{
		ApiID:       expected.apiID,
		ApiKey:      expected.apiKey,
		Environment: expected.environment,
		Version:     2,
	}

	if cfg.ApiID != expected.apiID {
		t.Errorf("Expected ApiID to be %s, but got %s", expected.apiID, cfg.ApiID)
	}

	if cfg.ApiKey != expected.apiKey {
		t.Errorf("Expected ApiKey to be %s, but got %s", expected.apiKey, cfg.ApiKey)
	}

	if cfg.Environment != expected.environment {
		t.Errorf("Expected URL Environment to be %s, but got %s", expected.environment, cfg.Environment)
	}

	if cfg.Version != expected.version {
		t.Errorf("Expected Version to be %d, but got %d", expected.version, cfg.Version)
	}
}

func TestEnvironments(t *testing.T) {
	expectedEnvironments := struct {
		sandbox    string
		production string
	}{
		sandbox:    `https://sandbox.ipag.com.br`,
		production: `https://api.ipag.com.br`,
	}

	if config.Environments.Sandbox != expectedEnvironments.sandbox {
		t.Errorf("Expected Sandbox URL Environment to be %s, but got %s", expectedEnvironments.sandbox, config.Environments.Sandbox)
	}

	if config.Environments.Production != expectedEnvironments.production {
		t.Errorf("Expected Production URL Environment to be %s, but got %s", expectedEnvironments.production, config.Environments.Production)
	}
}
