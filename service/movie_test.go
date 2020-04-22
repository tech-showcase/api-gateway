package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/api-gateway/model"
	"reflect"
	"testing"
)

func TestMovieService_Search(t *testing.T) {
	dummyEndpoints := model.MovieClientEndpoints{
		SearchMovieClientEndpoint: getDummyEndpoint(),
	}
	expectedOutput := getDummyResponse().MovieListPerPage

	ctx := context.Background()
	movieService := NewMovieService(dummyEndpoints)
	movies, err := movieService.Search(ctx, "Batman", 1)

	fmt.Println(movies)
	fmt.Println(expectedOutput)

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(movies, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func getDummyEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.SearchMovieRequest)

		if req.Keyword == "Batman" && req.PageNumber == 1 {
			return getDummyResponse(), nil
		}

		return model.SearchMovieResponse{}, errors.New("dummy error")
	}
}

func getDummyResponse() model.SearchMovieResponse {
	return model.SearchMovieResponse{
		MovieListPerPage: model.MovieListPerPage{
			Response:     "True",
			TotalResults: "375",
			Search: []model.MovieItem{
				{
					Title:  "Batman Beyond: Return of the Joker",
					Year:   "2000",
					ImdbID: "tt0233298",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BNmRmODEwNzctYzU1MS00ZDQ1LWI2NWMtZWFkNTQwNDg1ZDFiXkEyXkFqcGdeQXVyNTI4MjkwNjA@._V1_SX300.jpg",
				},
			},
		},
	}
}
