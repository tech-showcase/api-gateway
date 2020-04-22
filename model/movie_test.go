package model

import (
	"context"
	movieProto "github.com/tech-showcase/api-gateway/proto/movie"
	"reflect"
	"testing"
)

func TestEncodeSearchMovieGRPCRequest(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		dummyReq := SearchMovieRequest{
			Keyword:    "Batman",
			PageNumber: 1,
		}
		expectedReq := &movieProto.SearchRequest{
			Keyword:    dummyReq.Keyword,
			PageNumber: int32(dummyReq.PageNumber),
		}

		ctx := context.Background()
		reqAfterTranslation, err := encodeSearchMovieGRPCRequest(ctx, dummyReq)

		if err != nil {
			t.Fatal("an error has occurred")
		} else if !reflect.DeepEqual(expectedReq, reqAfterTranslation) {
			t.Fatal("unexpected output")
		}
	})
	t.Run("negative", func(t *testing.T) {
		dummyReq := SearchMovieResponse{}

		ctx := context.Background()
		_, err := encodeSearchMovieGRPCRequest(ctx, dummyReq)

		if err == nil {
			t.Fatal("an error should occur")
		}
	})
}

func TestDecodeSearchMovieGRPCResponse(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		dummyRes := &movieProto.SearchResponse{
			Response:     "True",
			TotalResults: "375",
			Search: []*movieProto.SearchResponse_MovieItem{
				{
					Title:  "Batman Beyond: Return of the Joker",
					Year:   "2000",
					ImdbID: "tt0233298",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BNmRmODEwNzctYzU1MS00ZDQ1LWI2NWMtZWFkNTQwNDg1ZDFiXkEyXkFqcGdeQXVyNTI4MjkwNjA@._V1_SX300.jpg",
				},
			},
		}
		expectedRes := SearchMovieResponse{
			MovieListPerPage: MovieListPerPage{
				Response:     dummyRes.Response,
				TotalResults: dummyRes.TotalResults,
				Search: []MovieItem{
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

		ctx := context.Background()
		resAfterTranslation, err := decodeSearchMovieGRPCResponse(ctx, dummyRes)

		if err != nil {
			t.Fatal("an error has occurred")
		} else if !reflect.DeepEqual(expectedRes, resAfterTranslation) {
			t.Fatal("unexpected output")
		}
	})
	t.Run("negative", func(t *testing.T) {
		dummyRes := SearchMovieRequest{}

		ctx := context.Background()
		_, err := decodeSearchMovieGRPCResponse(ctx, dummyRes)

		if err == nil {
			t.Fatal("an error should occur")
		}
	})
}
