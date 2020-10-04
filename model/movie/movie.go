package movie

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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

	instance.search = makeSearchMovieClientEndpoint(conn)

	return &instance, nil
}

func (instance *clientEndpoint) Search(ctx context.Context, req interface{}) (res interface{}, err error) {
	res, err = instance.search(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
