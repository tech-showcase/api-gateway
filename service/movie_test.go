package service

import (
	"context"
	"github.com/tech-showcase/api-gateway/model/movie"
	"reflect"
	"testing"
)

func TestMovieService_Search(t *testing.T) {
	expectedOutput := getDummyResponse()

	ctx := context.Background()
	dummyClientEndpoint := dummyClientEndpoint{}
	movieService := NewMovieService(&dummyClientEndpoint)

	req := movie.SearchMovieRequest{
		Keyword:    "Batman",
		PageNumber: 1,
	}
	movies, err := movieService.Search(ctx, req)

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(movies, expectedOutput) {
		t.Fatal("unexpected output")
	}
}
