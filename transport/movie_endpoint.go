package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/tech-showcase/api-gateway/model"
	"github.com/tech-showcase/api-gateway/presenter"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
)

type (
	MovieEndpoints struct {
		SearchMovieEndpoint endpoint.Endpoint
	}
)

func makeSearchMovieHTTPEndpoint(s service.MovieService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(presenter.SearchMovieRequest)
		result, err := s.Search(ctx, req.Keyword, req.PageNumber)
		res := presenter.SearchMovieResponse{
			SearchMovieResponse: model.SearchMovieResponse{
				MovieListPerPage: result,
			},
		}
		return res, nil
	}
}

func decodeSearchMovieHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req presenter.SearchMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
