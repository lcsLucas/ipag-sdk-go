package main

import (
	"github.com/lcslucas/ipag-sdk-go/credentials"
	"github.com/lcslucas/ipag-sdk-go/utils"
)

const (
	apiID   = "123456"
	apiKey  = "ABC123"
	version = 2
)

func main() {
	cfg := credentials.Credentials{
		ApiID:       apiID,
		ApiKey:      apiKey,
		Environment: credentials.Environments.Sandbox,
		Version:     version,
	}

	utils.PrintPretty(cfg)

}
