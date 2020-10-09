package covid19

import (
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/service"

	"time"
)

type (
	Endpoint struct {
		Get endpoint.HTTPEndpoint
	}
)

func NewCovid19Endpoint(covid19Services []service.Covid19Service, tracer stdopentracing.Tracer) (covid19Endpoint Endpoint) {
	endpointer := sd.FixedEndpointer{}
	for _, covid19Service := range covid19Services {
		getCovid19Endpoint := makeGetCovid19Endpoint(covid19Service)
		getCovid19Endpoint = middleware.ApplyTracerServer("getCovid19-endpoint", getCovid19Endpoint, tracer)

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
