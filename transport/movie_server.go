package transport

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/api-gateway/model"
	"github.com/tech-showcase/api-gateway/service"
	"io"
	"net/http"
)

func SearchMovieFactory() sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		movieClientEndpoints, conn := model.NewMovieGRPCClient(instance)
		movieService := service.NewMovieService(movieClientEndpoints)
		searchMovieEndpoint := MakeSearchMovieHTTPEndpoint(movieService)

		return searchMovieEndpoint, conn, nil
	}
}

func NewMovieHTTPServer(endpoints MovieEndpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/search", makeSearchMovieHTTPHandler(endpoints.SearchMovieEndpoint))

	return m
}

func makeSearchMovieHTTPHandler(searchMovieEndpoint endpoint.Endpoint) (handler *httptransport.Server) {
	handler = httptransport.NewServer(
		searchMovieEndpoint,
		decodeSearchMovieHTTPRequest,
		encodeHTTPResponse,
		[]httptransport.ServerOption{}...,
	)

	return
}
