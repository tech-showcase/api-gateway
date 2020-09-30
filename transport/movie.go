package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/endpoint/movie"
	"net/http"
)

func NewMovieHTTPServer(movieEndpoint movie.Endpoint) http.Handler {
	m := http.NewServeMux()
	m.Handle("/search", makeSearchMovieHTTPHandler(movieEndpoint.Search))

	return m
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
