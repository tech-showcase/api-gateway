package helper

import (
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

var TracerInstance stdopentracing.Tracer

func NewTracer() (stdopentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: "api-gateway",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "localhost:5775",
		},
	}

	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)

	stdopentracing.SetGlobalTracer(tracer)

	return tracer, closer, err
}
