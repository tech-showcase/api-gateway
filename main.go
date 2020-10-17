package main

import (
	"fmt"
	"github.com/tech-showcase/api-gateway/api"
	"github.com/tech-showcase/api-gateway/cmd"
	"github.com/tech-showcase/api-gateway/config"
	"github.com/tech-showcase/api-gateway/helper"
)

func init() {
	config.Instance = config.Read()

	//helper.LoggerInstance = helper.NewLogger()
	var err error
	helper.LoggerInstance, err = helper.NewFileLogger(config.Instance.Log.Filepath)
	if err != nil {
		panic(err)
	}

	helper.TracerInstance, _, err = helper.NewTracer(config.Instance.ServiceName, config.Instance.Tracer.AgentAddress)
	if err != nil {
		helper.LoggerInstance.Log("NewTracer", err)
	}

	helper.ConsulInstance, err = helper.NewConsulClient(config.Instance.Consul.AgentAddress)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hi, I am " + config.Instance.ServiceName + "!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
