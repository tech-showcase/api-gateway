package movie

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/helper"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
	"strconv"
	"time"
)

type (
	SearchMovieRequest struct {
		movie.SearchMovieRequest
	}
	SearchMovieResponse struct {
		movie.SearchMovieResponse
	}
)

func newSearchMovieConsulEndpoint(consulClient consulsd.Client, serviceName string, tracer stdopentracing.Tracer, logger log.Logger) endpoint.Endpoint {
	movieInstancer := consulsd.NewInstancer(consulClient, logger, serviceName, []string{}, true)
	searchMovieFactory := newSearchMovieFactory(movie.NewMovieClientEndpoint, logger, tracer)
	searchMovieEndpointer := sd.NewEndpointer(movieInstancer, searchMovieFactory, logger)

	balancer := lb.NewRoundRobin(searchMovieEndpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)

	return retry
}

func newSearchMovieFixedEndpoint(movieServices []service.MovieService, tracer stdopentracing.Tracer) endpoint.Endpoint {
	searchMovieEndpointer := sd.FixedEndpointer{}
	for _, movieService := range movieServices {
		searchMovieEndpoint := makeSearchMovieEndpoint(movieService)
		searchMovieEndpoint = middleware.ApplyTracerServer("searchMovie-endpoint", searchMovieEndpoint, tracer)

		searchMovieEndpointer = append(searchMovieEndpointer, searchMovieEndpoint)
	}

	balancer := lb.NewRoundRobin(searchMovieEndpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)

	return retry
}

func makeSearchMovieEndpoint(movieService service.MovieService) endpoint.Endpoint {
	searchMovieEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SearchMovieRequest)
		result, err := movieService.Search(ctx, req.SearchMovieRequest)
		if err != nil {
			return SearchMovieRequest{}, err
		}

		return SearchMovieResponse{SearchMovieResponse: result}, nil
	}

	return searchMovieEndpoint
}

func decodeSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pageNumberStr := helper.GetQueryStringValue(r, "page_number")
	pageNumber, _ := strconv.Atoi(pageNumberStr)

	req := SearchMovieRequest{
		SearchMovieRequest: movie.SearchMovieRequest{
			Keyword:    helper.GetQueryStringValue(r, "keyword"),
			PageNumber: pageNumber,
		},
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
