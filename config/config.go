package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type (
	Config struct {
		ConsulAddress               string `json:"consul_address"`
		EntertainmentServiceAddress string `json:"entertainment_service_address"`
	}
)

func Parse() (config Config, err error) {
	configPath := GetPath()

	configFileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}

	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		return
	}

	return
}

func GetPath() string {
	environment := "DEV"
	if environmentFromEnvVar := os.Getenv("ENVIRONMENT"); environmentFromEnvVar != "" {
		environment = environmentFromEnvVar
	}

	configPath := "config/config-dev.json"
	if configPathFromEnvVar := os.Getenv(environment + "_CONFIG_PATH"); configPathFromEnvVar != "" {
		configPath = configPathFromEnvVar
	}

	return configPath
}
