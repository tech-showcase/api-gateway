package api

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RegisterObservabilityHTTPAPI(r *mux.Router) {
	r.Handle("/observability/metrics", promhttp.Handler()).Methods(http.MethodGet)
}
