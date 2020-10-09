package movie

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/api-gateway/helper"
	"github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
	"strconv"
)

type (
	SearchMovieRequest struct {
		movie.SearchMovieRequest
	}
	SearchMovieResponse struct {
		movie.SearchMovieResponse
	}
)

func makeSearchMovieEndpoint(movieService service.MovieService) endpoint.Endpoint {
	searchMovieEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SearchMovieRequest)
		result, err := movieService.Search(ctx, req.SearchMovieRequest)
		if err != nil {
			return SearchMovieRequest{}, err
		}

		res := SearchMovieResponse{
			SearchMovieResponse: result,
		}

		return res, nil
	}

	return searchMovieEndpoint
}

func decodeSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	pageNumberStr := helper.GetQueryStringValue(r, "page_number")
	pageNumber, _ := strconv.Atoi(pageNumberStr)

	req := SearchMovieRequest{
		SearchMovieRequest: movie.SearchMovieRequest{
			Keyword:    helper.GetQueryStringValue(r, "keyword"),
			PageNumber: pageNumber,
		},
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
