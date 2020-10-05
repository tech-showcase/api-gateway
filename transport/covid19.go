package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/endpoint/covid19"
	"net/http"
)

func NewCovid19HTTPServer(covid19Endpoint covid19.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle("/covid19", makeGetCovid19HTTPHandler(covid19Endpoint.Get)).Methods(http.MethodGet)

	return r
}

func makeGetCovid19HTTPHandler(getCovid19Endpoint endpoint.HTTPEndpoint) (handler *httptransport.Server) {
	handler = httptransport.NewServer(
		getCovid19Endpoint.Endpoint,
		getCovid19Endpoint.Decoder,
		getCovid19Endpoint.Encoder,
		[]httptransport.ServerOption{}...,
	)

	return
}
