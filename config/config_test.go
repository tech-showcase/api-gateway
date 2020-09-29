package config

import (
	"os"
	"reflect"
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
	os.Setenv("ENTERTAINMENT_SERVICE_ADDRESS", dummyConfig.EntertainmentServiceAddress)
}

func getDummyConfig() Config {
	dummyConfig := Config{
		ConsulAddress:               "http://localhost",
		EntertainmentServiceAddress: "http://localhost",
	}

	return dummyConfig
}
