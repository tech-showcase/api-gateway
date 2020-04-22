package model

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	movieProto "github.com/tech-showcase/api-gateway/proto/movie"
	"google.golang.org/grpc"
)

type (
	MovieClientEndpoints struct {
		SearchMovieClientEndpoint endpoint.Endpoint
	}

	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		MovieListPerPage
	}
	MovieListPerPage struct {
		Response     string      `json:"Response"`
		Search       []MovieItem `json:"Search"`
		TotalResults string      `json:"totalResults"`
	}
	MovieItem struct {
		Poster string `json:"Poster"`
		Title  string `json:"Title"`
		Type   string `json:"Type"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
	}
)

func NewMovieGRPCClient(instance string) (movieEndpoints MovieClientEndpoints, conn *grpc.ClientConn) {
	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		return
	}

	searchMovieEndpoint := makeSearchMovieGRPCClient(conn)
	movieEndpoints = MovieClientEndpoints{
		SearchMovieClientEndpoint: searchMovieEndpoint,
	}

	return
}

func makeSearchMovieGRPCClient(conn *grpc.ClientConn) endpoint.Endpoint {
	clientOptions := make([]grpctransport.ClientOption, 0)
	searchMovieEndpoint := grpctransport.NewClient(
		conn,
		"Movie",
		"Search",
		encodeSearchMovieGRPCRequest,
		decodeSearchMovieGRPCResponse,
		movieProto.SearchResponse{},
		clientOptions...,
	).Endpoint()

	return searchMovieEndpoint
}

func encodeSearchMovieGRPCRequest(_ context.Context, r interface{}) (interface{}, error) {
	if res, ok := r.(SearchMovieRequest); ok {
		return &movieProto.SearchRequest{
			Keyword:    res.Keyword,
			PageNumber: int32(res.PageNumber),
		}, nil
	} else {
		return nil, errors.New("format request is wrong")
	}
}

func decodeSearchMovieGRPCResponse(_ context.Context, r interface{}) (interface{}, error) {
	if res, ok := r.(*movieProto.SearchResponse); ok {
		movies := make([]MovieItem, 0)
		for _, item := range res.Search {
			movie := MovieItem{
				Poster: item.Poster,
				Title:  item.Title,
				Type:   item.Type,
				Year:   item.Year,
				ImdbID: item.ImdbID,
			}
			movies = append(movies, movie)
		}

		return SearchMovieResponse{
			MovieListPerPage: MovieListPerPage{
				Response:     res.Response,
				Search:       movies,
				TotalResults: res.TotalResults,
			},
		}, nil
	} else {
		return nil, errors.New("format response is wrong")
	}
}
