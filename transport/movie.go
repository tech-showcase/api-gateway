package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/endpoint/movie"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMovieHTTPServer(movieEndpoint movie.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle("/movie", makeSearchMovieHTTPHandler(movieEndpoint.Search)).Methods(http.MethodGet)

	return r
}

func makeSearchMovieHTTPHandler(searchMovieEndpoint endpoint.HTTPEndpoint) (handler *httptransport.Server) {
	handler = httptransport.NewServer(
		searchMovieEndpoint.Endpoint,
		searchMovieEndpoint.Decoder,
		searchMovieEndpoint.Encoder,
		[]httptransport.ServerOption{}...,
	)

	return
}
