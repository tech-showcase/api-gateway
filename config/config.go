package config

import (
	"os"
	"strings"
)

type (
	Config struct {
		ConsulAddress                 string   `json:"consul_address"`
		EntertainmentServiceAddresses []string `json:"entertainment_service_address"`
	}
)

const (
	arrayDelimiter = ","
)

var Instance = Config{}

func Read() (config Config) {
	config = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config) {
	config.ConsulAddress = readEnvVarWithDefaultValue("CONSUL_ADDRESS", "http://localhost")
	config.EntertainmentServiceAddresses = readEnvVarArrayWithDefaultValue("ENTERTAINMENT_SERVICE_ADDRESS", []string{"localhost"})

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}

func readEnvVarArrayWithDefaultValue(key string, defaultValue []string) []string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return strings.Split(envVarValue, arrayDelimiter)
	}
	return defaultValue
}
