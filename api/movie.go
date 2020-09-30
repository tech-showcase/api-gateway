package api

import (
	"github.com/gorilla/mux"
	"github.com/tech-showcase/api-gateway/config"
	endpoint "github.com/tech-showcase/api-gateway/endpoint/movie"
	model "github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"github.com/tech-showcase/api-gateway/transport"
	"net/http"
)

func RegisterMovieHTTPAPI(r *mux.Router) {
	configInstance := config.Instance
	movieClientEndpoint := model.NewMovieClientEndpoint(configInstance.EntertainmentServiceAddress)
	movieService := service.NewMovieService(movieClientEndpoint)
	movieEndpoint := endpoint.NewMovieEndpoint(movieService)
	movieServer := transport.NewMovieHTTPServer(movieEndpoint)
	r.PathPrefix("/movie").Handler(http.StripPrefix("/movie", movieServer))
}
