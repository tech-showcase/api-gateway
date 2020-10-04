package movie

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/model/movie"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
)

type (
	SearchMovieRequest struct {
		movie.SearchMovieRequest
	}
	SearchMovieResponse struct {
		movie.SearchMovieResponse
	}
)

func makeSearchMovieEndpoint(movieService service.MovieService, logger log.Logger) endpoint.Endpoint {
	searchMovieEndpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SearchMovieRequest)
		result, err := movieService.Search(ctx, req.Keyword, req.PageNumber)
		if err != nil {
			return SearchMovieRequest{}, err
		}

		res := SearchMovieResponse{
			SearchMovieResponse: movie.SearchMovieResponse{
				ListPerPage: result,
			},
		}
		return res, nil
	}

	searchMovieEndpoint = middleware.ApplyCircuitBreaker("searchMovie", searchMovieEndpoint, logger)

	return searchMovieEndpoint
}

func decodeSearchMovieRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SearchMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
