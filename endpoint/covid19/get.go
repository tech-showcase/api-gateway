package covid19

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tech-showcase/api-gateway/helper"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/model/covid19"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
)

type (
	GetCovid19Request struct {
		covid19.GetCovid19Request
	}
	GetCovid19Response struct {
		covid19.GetCovid19Response
	}
)

func makeGetCovid19Endpoint(covid19Service service.Covid19Service, logger log.Logger) endpoint.Endpoint {
	getCovid19Endpoint := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCovid19Request)
		result, err := covid19Service.Get(ctx, req.GetCovid19Request)
		if err != nil {
			return GetCovid19Response{}, err
		}

		return GetCovid19Response{GetCovid19Response: result}, nil
	}

	getCovid19Endpoint = middleware.ApplyCircuitBreaker("getCovid19", getCovid19Endpoint, logger)

	return getCovid19Endpoint
}

func decodeGetCovid19Request(_ context.Context, r *http.Request) (interface{}, error) {
	fromStr := helper.GetQueryStringValue(r, "from")
	from, _ := helper.ParseDateTime(fromStr)

	toStr := helper.GetQueryStringValue(r, "to")
	to, _ := helper.ParseDateTime(toStr)

	req := GetCovid19Request{
		GetCovid19Request: covid19.GetCovid19Request{
			Country: helper.GetQueryStringValue(r, "country"),
			Status:  helper.GetQueryStringValue(r, "status"),
			From:    from,
			To:      to,
		},
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
