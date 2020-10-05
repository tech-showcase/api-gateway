package covid19

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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

func NewCovid19ClientEndpoint(covid19ServiceAddress string) (ClientEndpoint, error) {
	instance := clientEndpoint{}

	if !strings.HasPrefix(covid19ServiceAddress, "http") {
		covid19ServiceAddress = "http://" + covid19ServiceAddress
	}

	u, err := url.Parse(covid19ServiceAddress)
	if err != nil {
		return nil, err
	}
	instance.address = u

	instance.get = makeGetCovid19ClientEndpoint(u)

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
