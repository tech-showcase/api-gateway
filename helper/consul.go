package helper

import (
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

var ConsulInstance consulsd.Client

func NewConsulClient(agentAddress string) (consulsd.Client, error) {
	consulConfig := api.DefaultConfig()
	if len(agentAddress) > 0 {
		consulConfig.Address = agentAddress
	}

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}

	client := consulsd.NewClient(consulClient)
	return client, nil
}
