package middleware

import (
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/sony/gobreaker"
	"time"
)

func ApplyCircuitBreaker(name string, endpoint endpoint.Endpoint, logger log.Logger) (wrappedEndpoint endpoint.Endpoint) {
	wrappedEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    name,
		Timeout: 30 * time.Second,
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			logger.Log("ApplyCircuitBreaker", name, "from", from, "to", to)
		},
	}))(endpoint)

	return
}
