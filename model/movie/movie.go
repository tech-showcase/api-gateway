package movie

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/middleware"
	"google.golang.org/grpc"
)

type (
	ListPerPage struct {
		Response     string `json:"Response"`
		Search       []Item `json:"Search"`
		TotalResults string `json:"totalResults"`
	}
	Item struct {
		Poster string `json:"Poster"`
		Title  string `json:"Title"`
		Type   string `json:"Type"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
	}

	clientEndpoint struct {
		conn   *grpc.ClientConn
		search endpoint.Endpoint
	}
	ClientEndpoint interface {
		Search(context.Context, SearchMovieRequest) (SearchMovieResponse, error)
	}

	ModelFactory func(entertainmentServiceAddress string, logger log.Logger, tracer stdopentracing.Tracer) (ClientEndpoint, error)
)

func NewMovieClientEndpoint(entertainmentServiceAddress string, logger log.Logger, tracer stdopentracing.Tracer) (ClientEndpoint, error) {
	instance := clientEndpoint{}

	conn, err := grpc.Dial(entertainmentServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	instance.conn = conn

	searchMovieEndpoint := makeSearchMovieClientEndpoint(conn, logger, tracer)
	searchMovieEndpoint = middleware.ApplyTracerClient("searchMovie-model", searchMovieEndpoint, tracer)
	searchMovieEndpoint = middleware.ApplyLogger("searchMovie", searchMovieEndpoint, logger)
	searchMovieEndpoint = middleware.ApplyCircuitBreaker("searchMovie", searchMovieEndpoint, logger)
	instance.search = searchMovieEndpoint

	return &instance, nil
}

func (instance *clientEndpoint) Search(ctx context.Context, req SearchMovieRequest) (res SearchMovieResponse, err error) {
	response, err := instance.search(ctx, req)
	if err != nil {
		return SearchMovieResponse{}, err
	}

	res = response.(SearchMovieResponse)
	return res, nil
}
