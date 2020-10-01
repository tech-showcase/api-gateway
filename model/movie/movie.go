package movie

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	movieProto "github.com/tech-showcase/api-gateway/proto/movie"
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
		conn *grpc.ClientConn
	}
	ClientEndpoint interface {
		Search(context.Context, interface{}) (interface{}, error)
	}
)

func NewMovieClientEndpoint(entertainmentServiceAddress string) (ClientEndpoint, error) {
	instance := clientEndpoint{}

	conn, err := grpc.Dial(entertainmentServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	instance.conn = conn

	return &instance, nil
}

func (instance *clientEndpoint) Search(ctx context.Context, req interface{}) (res interface{}, err error) {
	clientOptions := make([]grpctransport.ClientOption, 0)
	searchMovieEndpoint := grpctransport.NewClient(
		instance.conn,
		"Movie",
		"Search",
		encodeSearchMovieRequest,
		decodeSearchMovieResponse,
		movieProto.SearchResponse{},
		clientOptions...,
	).Endpoint()

	res, err = searchMovieEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
