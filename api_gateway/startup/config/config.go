package config

import "os"

type Config struct {
	Port                string
	UserHost            string
	UserPort            string
	AutentificationHost string
	AutentificationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("GATEWAY_PORT"),
		UserHost:            os.Getenv("USER_SERVICE_HOST"),
		UserPort:            os.Getenv("USER_SERVICE_PORT"),
		AutentificationHost: os.Getenv("AUTENTIFICATION_SERVICE_HOST"),
		AutentificationPort: os.Getenv("AUTENTIFICATION_SERVICE_PORT"),
	}
}
