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
		Search(context.Context, string, int) (movie.ListPerPage, error)
	}
)

func NewMovieService(movieClientEndpoint movie.ClientEndpoint) MovieService {
	instance := movieService{}
	instance.movieClientEndpoint = movieClientEndpoint

	return &instance
}

func (instance *movieService) Search(ctx context.Context, keyword string, pageNumber int) (movies movie.ListPerPage, err error) {
	request := movie.SearchMovieRequest{
		Keyword:    keyword,
		PageNumber: pageNumber,
	}
	response, err := instance.movieClientEndpoint.Search(ctx, request)
	if err != nil {
		return
	}

	movies = response.ListPerPage
	return
}
