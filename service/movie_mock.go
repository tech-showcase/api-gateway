package service

import (
	"context"
	"errors"
	"github.com/tech-showcase/api-gateway/model/movie"
)

type (
	dummyClientEndpoint struct{}
)

func (instance *dummyClientEndpoint) Search(ctx context.Context, req movie.SearchMovieRequest) (res movie.SearchMovieResponse, err error) {
	if req.Keyword == "Batman" && req.PageNumber == 1 {
		return getDummyResponse(), nil
	}

	return movie.SearchMovieResponse{}, errors.New("dummy error")
}

func getDummyResponse() movie.SearchMovieResponse {
	return movie.SearchMovieResponse{
		ListPerPage: movie.ListPerPage{
			Response:     "True",
			TotalResults: "375",
			Search: []movie.Item{
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
