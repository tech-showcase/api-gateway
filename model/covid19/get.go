package covid19

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/helper"
	"net/http"
	"net/url"
	"time"
)

type (
	GetCovid19Request struct {
		Country string    `json:"country"`
		Status  string    `json:"status"`
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
	}
	GetCovid19Response struct {
		Data `json:"data"`
	}
)

func makeGetCovid19ClientEndpoint(covid19ServiceURL *url.URL, logger log.Logger, tracer stdopentracing.Tracer) endpoint.Endpoint {
	getCovid19URL, _ := helper.JoinURL(covid19ServiceURL, "/covid19")

	clientOptions := make([]httptransport.ClientOption, 0)
	clientOptions = append(clientOptions, httptransport.ClientBefore(opentracing.ContextToHTTP(tracer, logger)))

	getCovid19Endpoint := httptransport.NewClient(
		http.MethodGet,
		getCovid19URL,
		encodeGetCovid19HTTPRequest,
		decodeGetCovid19HTTPResponse,
		clientOptions...,
	).Endpoint()

	return getCovid19Endpoint
}

func encodeGetCovid19HTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	if req, ok := request.(GetCovid19Request); ok {
		q := r.URL.Query()
		q.Add("country", req.Country)
		q.Add("status", req.Status)
		q.Add("from", req.From.Format(time.RFC3339Nano))
		q.Add("to", req.To.Format(time.RFC3339Nano))
		r.URL.RawQuery = q.Encode()

		return nil
	} else {
		return errors.New("request format is wrong")
	}
}

func decodeGetCovid19HTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var req GetCovid19Response
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
