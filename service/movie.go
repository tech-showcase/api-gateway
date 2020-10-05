package service

import (
	"context"
	"github.com/tech-showcase/api-gateway/model/movie"
)

type (
	movieService struct {
		movieClientEndpoint movie.ClientEndpoint
	}
	MovieService interface {
		Search(context.Context, movie.SearchMovieRequest) (movie.SearchMovieResponse, error)
	}
)

func NewMovieService(movieClientEndpoint movie.ClientEndpoint) MovieService {
	instance := movieService{}
	instance.movieClientEndpoint = movieClientEndpoint

	return &instance
}

func (instance *movieService) Search(ctx context.Context, req movie.SearchMovieRequest) (res movie.SearchMovieResponse, err error) {
	res, err = instance.movieClientEndpoint.Search(ctx, req)
	if err != nil {
		return
	}

	return
}
