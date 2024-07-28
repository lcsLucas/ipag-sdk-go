package config_test

import (
	"testing"
	"time"

	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/credentials"
)

func TestConfig(t *testing.T) {
	expectedClientConfig := struct {
		timeout     time.Duration
		readTimeout time.Duration
		headers     map[string]string
	}{
		timeout:     30 * time.Second,
		readTimeout: 30 * time.Second,
		headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}

	expectedCredentialsConfig := struct {
		apiID       string
		apiKey      string
		environment string
		version     uint8
	}{
		apiID:       "apiID",
		apiKey:      "apiKey",
		environment: "environment",
		version:     1,
	}

	expectedConfig := struct {
		credentials struct {
			apiID       string
			apiKey      string
			environment string
			version     uint8
		}
		client struct {
			timeout     time.Duration
			readTimeout time.Duration
			headers     map[string]string
		}
	}{
		client:      expectedClientConfig,
		credentials: expectedCredentialsConfig,
	}

	cli := config.Config{
		Credentials: credentials.Credentials{
			ApiID:       expectedConfig.credentials.apiID,
			ApiKey:      expectedConfig.credentials.apiKey,
			Environment: expectedConfig.credentials.environment,
			Version:     expectedConfig.credentials.version,
		},
		Client: config.ClientConfig{
			Timeout:     expectedConfig.client.timeout,
			ReadTimeout: expectedConfig.client.readTimeout,
			Headers:     expectedConfig.client.headers,
		},
	}

	if cli.Credentials.ApiID != expectedConfig.credentials.apiID {
		t.Errorf("Expected %s, got %s", expectedConfig.credentials.apiID, cli.Credentials.ApiID)
	}

	if cli.Credentials.ApiKey != expectedConfig.credentials.apiKey {
		t.Errorf("Expected %s, got %s", expectedConfig.credentials.apiKey, cli.Credentials.ApiKey)
	}

	if cli.Credentials.Environment != expectedConfig.credentials.environment {
		t.Errorf("Expected %s, got %s", expectedConfig.credentials.environment, cli.Credentials.Environment)
	}

	if cli.Credentials.Version != expectedConfig.credentials.version {
		t.Errorf("Expected %d, got %d", expectedConfig.credentials.version, cli.Credentials.Version)
	}

	if cli.Client.Timeout != expectedConfig.client.timeout {
		t.Errorf("Expected %v, got %v", expectedConfig.client.timeout, cli.Client.Timeout)
	}

	if cli.Client.ReadTimeout != expectedConfig.client.readTimeout {
		t.Errorf("Expected %v, got %v", expectedConfig.client.readTimeout, cli.Client.ReadTimeout)
	}

	if len(cli.Client.Headers) != len(expectedConfig.client.headers) {
		t.Errorf("Expected %d, got %d", len(expectedConfig.client.headers), len(cli.Client.Headers))
	}

}
