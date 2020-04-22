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

