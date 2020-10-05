package covid19

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/service"

	"time"
)

type (
	Endpoint struct {
		Get endpoint.HTTPEndpoint
	}
)

func NewCovid19Endpoint(covid19Services []service.Covid19Service, logger log.Logger) (covid19Endpoint Endpoint) {
	endpointer := sd.FixedEndpointer{}
	for _, covid19Service := range covid19Services {
		getCovid19Endpoint := makeGetCovid19Endpoint(covid19Service, logger)
		endpointer = append(endpointer, getCovid19Endpoint)
	}

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)

	covid19Endpoint.Get = endpoint.HTTPEndpoint{
		Endpoint: retry,
		Decoder:  decodeGetCovid19Request,
		Encoder:  encodeResponse,
	}

	return
}
