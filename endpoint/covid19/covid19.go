package covid19

import (
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/service"
)

type (
	Endpoint struct {
		Get endpoint.HTTPEndpoint
	}
)

func NewCovid19Endpoint(covid19Services []service.Covid19Service, tracer stdopentracing.Tracer) (covid19Endpoint Endpoint) {
	getCovid19Endpoint := newGetCovid19FixedEndpoint(covid19Services, tracer)

	covid19Endpoint.Get = endpoint.HTTPEndpoint{
		Endpoint: getCovid19Endpoint,
		Decoder:  decodeGetCovid19Request,
		Encoder:  encodeResponse,
	}

	return
}
