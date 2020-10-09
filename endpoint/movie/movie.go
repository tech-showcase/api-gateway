package movie

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
		Search endpoint.HTTPEndpoint
	}
)

func NewMovieEndpoint(movieServices []service.MovieService, tracer stdopentracing.Tracer) (movieEndpoint Endpoint) {
	endpointer := sd.FixedEndpointer{}
	for _, movieService := range movieServices {
		searchMovieEndpoint := makeSearchMovieEndpoint(movieService)
		searchMovieEndpoint = middleware.ApplyTracerServer("searchMovie-endpoint", searchMovieEndpoint, tracer)

		endpointer = append(endpointer, searchMovieEndpoint)
	}

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)

	movieEndpoint.Search = endpoint.HTTPEndpoint{
		Endpoint: retry,
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeResponse,
	}

	return movieEndpoint
}
