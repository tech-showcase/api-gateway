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

	os.Setenv("SERVICE_NAME", dummyConfig.ServiceName)
	os.Setenv("CONSUL_ADDRESS", dummyConfig.ConsulAddress)
	os.Setenv("ENTERTAINMENT_SERVICE_ADDRESS", strings.Join(dummyConfig.EntertainmentServiceAddresses, arrayDelimiter))
	os.Setenv("COVID19_SERVICE_ADDRESS", strings.Join(dummyConfig.Covid19ServiceAddresses, arrayDelimiter))
	os.Setenv("TRACER_AGENT_ADDRESS", dummyConfig.Tracer.AgentAddress)
}

func getDummyConfig() Config {
	dummyConfig := Config{
		ServiceName:   "service-name",
		ConsulAddress: "consul-address",
		EntertainmentServiceAddresses: []string{
			"entertainment-service-address-1",
			"entertainment-service-address-2",
		},
		Covid19ServiceAddresses: []string{
			"covid19-service-address-1",
			"covid19-service-address-2",
		},
		Tracer: Tracer{
			AgentAddress: "tracer-agent-address",
		},
	}

	return dummyConfig
}
