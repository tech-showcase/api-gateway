package service

import (
	"context"
	"github.com/tech-showcase/api-gateway/model"
)

type (
	movieService struct {
		movieClientEndpoints model.MovieClientEndpoints
	}
	MovieService interface {
		Search(context.Context, string, int) (model.MovieListPerPage, error)
	}
)

func NewMovieService(movieClientEndpoints model.MovieClientEndpoints) MovieService {
	instance := movieService{}
	instance.movieClientEndpoints = movieClientEndpoints

	return &instance
}

func (instance *movieService) Search(ctx context.Context, keyword string, pageNumber int) (movies model.MovieListPerPage, err error) {
	request := model.SearchMovieRequest{
		Keyword:    keyword,
		PageNumber: pageNumber,
	}
	response, err := instance.movieClientEndpoints.SearchMovieClientEndpoint(ctx, request)
	if err != nil {
		return
	}

	searchMovieResponse := response.(model.SearchMovieResponse)
	movies = searchMovieResponse.MovieListPerPage

	return
}
