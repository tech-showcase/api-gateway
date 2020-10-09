package covid19

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/tech-showcase/api-gateway/middleware"
	"net/url"
	"strings"
	"time"
)

type (
	Data []struct {
		Country     string    `json:"Country"`
		CountryCode string    `json:"CountryCode"`
		Province    string    `json:"Province"`
		City        string    `json:"City"`
		CityCode    string    `json:"CityCode"`
		Lat         string    `json:"Lat"`
		Lon         string    `json:"Lon"`
		Cases       int       `json:"Cases"`
		Status      string    `json:"Status"`
		Date        time.Time `json:"Date"`
	}

	clientEndpoint struct {
		address *url.URL
		get     endpoint.Endpoint
	}
	ClientEndpoint interface {
		Get(context.Context, GetCovid19Request) (GetCovid19Response, error)
	}
)

func NewCovid19ClientEndpoint(covid19ServiceAddress string, logger log.Logger, tracer stdopentracing.Tracer) (ClientEndpoint, error) {
	instance := clientEndpoint{}

	if !strings.HasPrefix(covid19ServiceAddress, "http") {
		covid19ServiceAddress = "http://" + covid19ServiceAddress
	}

	u, err := url.Parse(covid19ServiceAddress)
	if err != nil {
		return nil, err
	}
	instance.address = u

	getCovid19Endpoint := makeGetCovid19ClientEndpoint(u, logger, tracer)
	getCovid19Endpoint = middleware.ApplyTracerClient("getCovid19-model", getCovid19Endpoint, tracer)
	getCovid19Endpoint = middleware.ApplyCircuitBreaker("getCovid19", getCovid19Endpoint, logger)
	instance.get = getCovid19Endpoint

	return &instance, nil
}

func (instance *clientEndpoint) Get(ctx context.Context, req GetCovid19Request) (res GetCovid19Response, err error) {
	response, err := instance.get(ctx, req)
	if err != nil {
		return GetCovid19Response{}, err
	}

	res = response.(GetCovid19Response)
	return res, nil
}
