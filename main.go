package main

import (
	"fmt"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/gorilla/mux"
	"github.com/tech-showcase/api-gateway/cmd"
	"github.com/tech-showcase/api-gateway/global"
	"github.com/tech-showcase/api-gateway/helper"
	"github.com/tech-showcase/api-gateway/transport"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hi, I am API Gateway!")

	args := cmd.Parse()
	config := global.Configuration

	consulClient, err := helper.NewConsul(config.ConsulAddress)
	if err != nil {
		panic(err)
	}

	movieEndpoints := transport.MovieEndpoints{}
	logger := helper.NewLogger()

	factory := transport.SearchMovieFactory()
	instancer := consulsd.NewInstancer(consulClient, logger, "movie", []string{}, true)
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)
	movieEndpoints.SearchMovieEndpoint = retry

	r := mux.NewRouter()
	movieServer := transport.NewMovieHTTPServer(movieEndpoints)
	r.PathPrefix("/movie").Handler(http.StripPrefix("/movie", movieServer))

	portStr := fmt.Sprintf(":%d", args.Port)
	http.ListenAndServe(portStr, r)
}
