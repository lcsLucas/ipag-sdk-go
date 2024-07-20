package main

import (
	"github.com/lcslucas/ipag-sdk-go/config"
	"github.com/lcslucas/ipag-sdk-go/utils"
)

const (
	apiID   = "123456"
	apiKey  = "ABC123"
	version = 2
)

func main() {
	cfg := config.Config{
		ApiID:       apiID,
		ApiKey:      apiKey,
		Environment: config.Environments.Sandbox,
		Version:     version,
	}

	utils.PrintPretty(cfg)

}
