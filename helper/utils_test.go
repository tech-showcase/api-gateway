package helper

import (
	"net/url"
	"testing"
)

func TestJoinURL(t *testing.T) {
	dummyBaseURL, _ := url.Parse("http://localhost:8080/")
	dummyPath := "/v1/resources"
	expectedURL := "http://localhost:8080/v1/resources"

	joinedURL, err := JoinURL(dummyBaseURL, dummyPath)
	if err != nil {
		t.Fatal("an error has occurred")
		return
	}

	urlStr := joinedURL.String()
	if urlStr != expectedURL {
		t.Fatal("unexpected output")
	}
}
