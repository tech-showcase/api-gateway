package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/tech-showcase/api-gateway/presenter"
	"net/http"
	"reflect"
	"testing"
)

func TestDecodeSearchMovieHTTPRequest(t *testing.T) {
	dummyReq := getDummyRequest()
	expectedReq := getExpectedRequest()

	ctx := context.Background()
	reqAfterTranslation, err := decodeSearchMovieHTTPRequest(ctx, dummyReq)

	if err != nil {
		t.Fatal("an error has occurred")
	} else if !reflect.DeepEqual(reqAfterTranslation, expectedReq) {
		t.Fatal("unexpected output")
	}
}

func getDummyRequest() *http.Request {
	jsonBody := getDummyRequestBody()

	req, err := http.NewRequest("POST", "http://localhost", bytes.NewBufferString(jsonBody))
	if err != nil {
		return nil
	}

	return req
}

func getExpectedRequest() (req presenter.SearchMovieRequest) {
	jsonBody := getDummyRequestBody()
	json.NewDecoder(bytes.NewBufferString(jsonBody)).Decode(&req)

	return
}

func getDummyRequestBody() string {
	jsonBody := `{
		"keyword": "Batman",
		"page_number": 1
		}`
	return jsonBody
}
