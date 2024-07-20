package config

type Config struct {
	ApiID       string
	ApiKey      string
	Environment string
	Version     uint8
}

type environments struct {
	Sandbox    string
	Production string
}

var Environments = environments{
	Sandbox:    `https://sandbox.ipag.com.br`,
	Production: `https://api.ipag.com.br`,
}
