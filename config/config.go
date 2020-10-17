package config

import (
	"os"
	"strings"
)

type (
	Config struct {
		ServiceName                   string   `json:"service_name"`
		EntertainmentServiceAddresses []string `json:"entertainment_service_address"`
		Covid19ServiceAddresses       []string `json:"covid19_service_address"`
		Tracer                        Tracer   `json:"tracer"`
		Consul                        Consul   `json:"consul"`
		Log                           Log      `json:"log"`
	}

	Tracer struct {
		AgentAddress string `json:"agent_address"`
	}

	Consul struct {
		AgentAddress string `json:"agent_address"`
	}

	Log struct {
		Filepath string `json:"filepath"`
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
	config.ServiceName = readEnvVarWithDefaultValue("SERVICE_NAME", "api-gateway")
	config.EntertainmentServiceAddresses = readEnvVarArrayWithDefaultValue("ENTERTAINMENT_SERVICE_ADDRESS", []string{"localhost:8085"})
	config.Covid19ServiceAddresses = readEnvVarArrayWithDefaultValue("COVID19_SERVICE_ADDRESS", []string{"localhost:8083"})
	config.Tracer.AgentAddress = readEnvVarWithDefaultValue("TRACER_AGENT_ADDRESS", "localhost:5775")
	config.Consul.AgentAddress = readEnvVarWithDefaultValue("CONSUL_AGENT_ADDRESS", "localhost:8500")
	config.Log.Filepath = readEnvVarWithDefaultValue("LOG_FILEPATH", "./server.log")

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
