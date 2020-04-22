package transport

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

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
