package movie

import (
	"github.com/go-kit/kit/log"
	consulsd "github.com/go-kit/kit/sd/consul"
	stdopentracing "github.com/opentracing/opentracing-go"
	generalEndpoint "github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/service"
)

type (
	Endpoint struct {
		Search generalEndpoint.HTTPEndpoint
	}
)

func NewMovieEndpoint(movieServices []service.MovieService, tracer stdopentracing.Tracer) (movieEndpoint Endpoint) {
	searchMovieEndpoint := newSearchMovieFixedEndpoint(movieServices, tracer)

	movieEndpoint.Search = generalEndpoint.HTTPEndpoint{
		Endpoint: searchMovieEndpoint,
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeResponse,
	}

	return movieEndpoint
}

func NewConsulMovieEndpoint(consulClient consulsd.Client, tracer stdopentracing.Tracer, logger log.Logger) (movieEndpoint Endpoint) {
	searchMovieEndpoint := newSearchMovieConsulEndpoint(consulClient, "entertainment-service", tracer, logger)

	movieEndpoint.Search = generalEndpoint.HTTPEndpoint{
		Endpoint: searchMovieEndpoint,
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeResponse,
	}

	return movieEndpoint
}
