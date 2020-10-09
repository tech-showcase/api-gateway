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

	helper.LoggerInstance = helper.NewLogger()

	var err error
	helper.TracerInstance, _, err = helper.NewTracer(config.Instance.ServiceName, config.Instance.Tracer.AgentAddress)
	if err != nil {
		helper.LoggerInstance.Log("NewTracer", err)
	}
}

func main() {
	fmt.Println("Hi, I am " + config.Instance.ServiceName + "!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
