package service

import (
	"context"
	"github.com/tech-showcase/api-gateway/model/covid19"
)

type (
	covid19Service struct {
		covid19ClientEndpoint covid19.ClientEndpoint
	}
	Covid19Service interface {
		Get(context.Context, covid19.GetCovid19Request) (covid19.GetCovid19Response, error)
	}
)

func NewCovid19Service(covid19ClientEndpoint covid19.ClientEndpoint) Covid19Service {
	instance := covid19Service{}
	instance.covid19ClientEndpoint = covid19ClientEndpoint

	return &instance
}

func (instance *covid19Service) Get(ctx context.Context, req covid19.GetCovid19Request) (res covid19.GetCovid19Response, err error) {
	response, err := instance.covid19ClientEndpoint.Get(ctx, req)
	if err != nil {
		return
	}

	res = response
	return
}
