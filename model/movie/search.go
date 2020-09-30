package movie

import (
	"context"
	"errors"
	movieProto "github.com/tech-showcase/api-gateway/proto/movie"
)

type (
	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		ListPerPage
	}
)

func encodeSearchMovieRequest(_ context.Context, r interface{}) (interface{}, error) {
	if req, ok := r.(SearchMovieRequest); ok {
		return &movieProto.SearchRequest{
			Keyword:    req.Keyword,
			PageNumber: int32(req.PageNumber),
		}, nil
	} else {
		return nil, errors.New("format request is wrong")
	}
}

func decodeSearchMovieResponse(_ context.Context, r interface{}) (interface{}, error) {
	if res, ok := r.(*movieProto.SearchResponse); ok {
		movies := make([]Item, 0)
		for _, item := range res.Search {
			movie := Item{
				Poster: item.Poster,
				Title:  item.Title,
				Type:   item.Type,
				Year:   item.Year,
				ImdbID: item.ImdbId,
			}
			movies = append(movies, movie)
		}

		return SearchMovieResponse{
			ListPerPage: ListPerPage{
				Response:     res.Response,
				Search:       movies,
				TotalResults: res.TotalResults,
			},
		}, nil
	} else {
		return nil, errors.New("format response is wrong")
	}
}