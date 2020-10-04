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
}

func main() {
	fmt.Println("Hi, I am API Gateway!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
