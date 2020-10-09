package api

import (
	"github.com/gorilla/mux"
	"github.com/tech-showcase/api-gateway/config"
	endpoint "github.com/tech-showcase/api-gateway/endpoint/movie"
	"github.com/tech-showcase/api-gateway/helper"
	model "github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"github.com/tech-showcase/api-gateway/transport"
)

func RegisterMovieHTTPAPI(r *mux.Router) {
	configInstance := config.Instance
	loggerInstance := helper.LoggerInstance
	tracerInstance := helper.TracerInstance

	var movieServices []service.MovieService
	for _, entertainmentServiceAddress := range configInstance.EntertainmentServiceAddresses {
		movieClientEndpoint, err := model.NewMovieClientEndpoint(entertainmentServiceAddress, loggerInstance, tracerInstance)
		if err != nil {
			panic(err)
		}
		movieService := service.NewMovieService(movieClientEndpoint)
		movieServices = append(movieServices, movieService)
	}

	movieEndpoint := endpoint.NewMovieEndpoint(movieServices, tracerInstance)
	movieServer := transport.NewMovieHTTPServer(movieEndpoint)
	r.PathPrefix("/movie").Handler(movieServer)
}
