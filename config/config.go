package config

import (
	"os"
)

type (
	Config struct {
		ConsulAddress               string `json:"consul_address"`
		EntertainmentServiceAddress string `json:"entertainment_service_address"`
	}
)

var Instance = Config{}

func Read() (config Config) {
	config = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config) {
	config.ConsulAddress = readEnvVarWithDefaultValue("CONSUL_ADDRESS", "http://localhost")
	config.EntertainmentServiceAddress = readEnvVarWithDefaultValue("ENTERTAINMENT_SERVICE_ADDRESS", "http://localhost")

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}
