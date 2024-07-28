package config

import (
	"time"

	"github.com/lcslucas/ipag-sdk-go/credentials"
)

type ClientConfig struct {
	Timeout     time.Duration
	ReadTimeout time.Duration
	Headers     map[string]string
}

type Config struct {
	Credentials credentials.Credentials
	Client      ClientConfig
}
