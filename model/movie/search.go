package movie

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	movieProto "github.com/tech-showcase/api-gateway/proto/movie"
	"google.golang.org/grpc"
)

type (
	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		ListPerPage
	}
)

func makeSearchMovieClientEndpoint(conn *grpc.ClientConn, logger log.Logger, tracer stdopentracing.Tracer) endpoint.Endpoint {
	clientOptions := make([]grpctransport.ClientOption, 0)
	clientOptions = append(clientOptions, grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)))

	searchMovieEndpoint := grpctransport.NewClient(
		conn,
		"Movie",
		"Search",
		encodeSearchMovieRequest,
		decodeSearchMovieResponse,
		movieProto.SearchResponse{},
		clientOptions...,
	).Endpoint()

	return searchMovieEndpoint
}

func encodeSearchMovieRequest(_ context.Context, r interface{}) (interface{}, error) {
	if req, ok := r.(SearchMovieRequest); ok {
		return &movieProto.SearchRequest{
			Keyword:    req.Keyword,
			PageNumber: int32(req.PageNumber),
		}, nil
	} else {
		return nil, errors.New("request format is wrong")
	}
}

func decodeSearchMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	if res, ok := r.(*movieProto.SearchResponse); ok {
		movies := make([]Item, 0)
		for _, item := range res.Search {
			movie := Item{
				Poster: item.Poster,
				Title:  item.Title,
				Type:   item.Type,
				Year:   item.Year,
				ImdbID: item.ImdbId,
			}
			movies = append(movies, movie)
		}

		return SearchMovieResponse{
			ListPerPage: ListPerPage{
				Response:     res.Response,
				Search:       movies,
				TotalResults: res.TotalResults,
			},
		}, nil
	} else {
		return nil, errors.New("response format is wrong")
	}
}
