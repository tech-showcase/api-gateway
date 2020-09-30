package service

import (
	"context"
	"reflect"
	"testing"
)

func TestMovieService_Search(t *testing.T) {
	expectedOutput := getDummyResponse().ListPerPage

	ctx := context.Background()
	dummyClientEndpoint := dummyClientEndpoint{}
	movieService := NewMovieService(&dummyClientEndpoint)
	movies, err := movieService.Search(ctx, "Batman", 1)

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(movies, expectedOutput) {
		t.Fatal("unexpected output")
	}
}
