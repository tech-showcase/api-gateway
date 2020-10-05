package api

import (
	"github.com/gorilla/mux"
	"github.com/tech-showcase/api-gateway/config"
	endpoint "github.com/tech-showcase/api-gateway/endpoint/covid19"
	"github.com/tech-showcase/api-gateway/helper"
	model "github.com/tech-showcase/api-gateway/model/covid19"
	"github.com/tech-showcase/api-gateway/service"
	"github.com/tech-showcase/api-gateway/transport"
)

func RegisterCovid19HTTPAPI(r *mux.Router) {
	configInstance := config.Instance
	loggerInstance := helper.LoggerInstance

	var covid19Services []service.Covid19Service
	for _, covid19ServiceAddress := range configInstance.Covid19ServiceAddresses {
		covid19ClientEndpoint, err := model.NewCovid19ClientEndpoint(covid19ServiceAddress)
		if err != nil {
			panic(err)
		}
		covid19Service := service.NewCovid19Service(covid19ClientEndpoint)
		covid19Services = append(covid19Services, covid19Service)
	}

	covid19Endpoint := endpoint.NewCovid19Endpoint(covid19Services, loggerInstance)
	covid19Server := transport.NewCovid19HTTPServer(covid19Endpoint)
	r.PathPrefix("/covid19").Handler(covid19Server)
}
