package covid19

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/tech-showcase/api-gateway/middleware"
	"github.com/tech-showcase/api-gateway/model/covid19"
	"github.com/tech-showcase/api-gateway/service"
	"net/http"
	"time"
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

		res := GetCovid19Response{
			GetCovid19Response: result,
		}
		return res, nil
	}

	getCovid19Endpoint = middleware.ApplyCircuitBreaker("getCovid19", getCovid19Endpoint, logger)

	return getCovid19Endpoint
}

func decodeGetCovid19Request(_ context.Context, r *http.Request) (interface{}, error) {
	fromStr := getQueryStringValue(r, "from")
	from, _ := parseDateTime(fromStr)

	toStr := getQueryStringValue(r, "to")
	to, _ := parseDateTime(toStr)

	req := GetCovid19Request{
		GetCovid19Request: covid19.GetCovid19Request{
			Country: getQueryStringValue(r, "country"),
			Status:  getQueryStringValue(r, "status"),
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

func getQueryStringValue(r *http.Request, key string) (value string) {
	if valueArr, ok := r.URL.Query()[key]; ok {
		value = valueArr[0]
	}

	return
}

func parseDateTime(timeStr string) (value time.Time, err error) {
	value, err = time.Parse(time.RFC3339Nano, timeStr)
	if err != nil {
		return
	}

	return
}
