package movie

import (
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/tech-showcase/api-gateway/endpoint"
	"github.com/tech-showcase/api-gateway/service"
	"time"
)

type (
	Endpoint struct {
		Search endpoint.HTTPEndpoint
	}
)

func NewMovieEndpoint(movieService service.MovieService) (movieEndpoint Endpoint) {
	endpointer := sd.FixedEndpointer{}
	searchMovieEndpoint := makeSearchMovieEndpoint(movieService)
	endpointer = append(endpointer, searchMovieEndpoint)

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(3, 500*time.Millisecond, balancer)

	movieEndpoint.Search = endpoint.HTTPEndpoint{
		Endpoint: retry,
		Decoder:  decodeSearchMovieRequest,
		Encoder:  encodeResponse,
	}

	return movieEndpoint
}
