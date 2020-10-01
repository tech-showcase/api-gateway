package config

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	setDummyEnvVar()
	expectedOutput := getDummyConfig()

	config := Read()

	if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func setDummyEnvVar() {
	dummyConfig := getDummyConfig()

	os.Setenv("CONSUL_ADDRESS", dummyConfig.ConsulAddress)
	os.Setenv("ENTERTAINMENT_SERVICE_ADDRESS", strings.Join(dummyConfig.EntertainmentServiceAddresses, arrayDelimiter))
}

func getDummyConfig() Config {
	dummyConfig := Config{
		ConsulAddress: "http://localhost",
		EntertainmentServiceAddresses: []string{
			"localhost",
			"127.0.0.1",
		},
	}

	return dummyConfig
}
