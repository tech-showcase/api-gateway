package movie

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"io"
)

func newSearchMovieFactory(makeModel movie.ModelFactory, logger log.Logger, tracer stdopentracing.Tracer) sd.Factory {
	return func(entertainmentServiceAddress string) (endpoint.Endpoint, io.Closer, error) {
		movieClientEndpoint, err := makeModel(entertainmentServiceAddress, logger, tracer)
		if err != nil {
			return nil, nil, err
		}

		movieService := service.NewMovieService(movieClientEndpoint)

		searchMovieEndpoint := makeSearchMovieEndpoint(movieService)
		searchMovieEndpoint = middleware.ApplyTracerServer("searchMovie-endpoint", searchMovieEndpoint, tracer)

		return searchMovieEndpoint, nil, nil
	}
}
