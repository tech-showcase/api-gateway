package middleware

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
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

func ApplyTracerServer(operationName string, endpoint endpoint.Endpoint, tracer stdopentracing.Tracer) (wrappedEndpoint endpoint.Endpoint) {
	wrappedEndpoint = opentracing.TraceServer(tracer, operationName)(endpoint)
	return
}

func ApplyTracerClient(operationName string, endpoint endpoint.Endpoint, tracer stdopentracing.Tracer) (wrappedEndpoint endpoint.Endpoint) {
	wrappedEndpoint = opentracing.TraceClient(tracer, operationName)(endpoint)
	return
}

func ApplyLogger(operationName string, nextEndpoint endpoint.Endpoint, logger log.Logger) (wrappedEndpoint endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		defer func(begin time.Time) {
			logger.Log("operationName", operationName, "elapsedTime", time.Since(begin))
		}(time.Now())

		return nextEndpoint(ctx, request)
	}
}
