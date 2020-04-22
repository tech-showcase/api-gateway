package helper

import (
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

func NewConsul(consulAddress string) (client consulsd.Client, err error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulAddress

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return
	}

	client = consulsd.NewClient(consulClient)

	return
}
