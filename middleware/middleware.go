package middleware

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"github.com/tech-showcase/api-gateway/config"
	"strings"
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

func ApplyMetrics(subsystem, operation string, nextEndpoint endpoint.Endpoint) endpoint.Endpoint {
	var labelNames []string
	namespace := strings.Replace(config.Instance.ServiceName, "-", "_", -1)

	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      operation + "_request_count",
		Help:      "Number of requests received.",
	}, labelNames)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      operation + "_request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, labelNames)

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		requestCount.Add(1)
		defer func(begin time.Time) {
			requestLatency.Observe(time.Since(begin).Seconds())
		}(time.Now())

		return nextEndpoint(ctx, request)
	}
}
