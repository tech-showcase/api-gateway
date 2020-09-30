package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type (
	HTTPEndpoint struct {
		Endpoint endpoint.Endpoint
		Decoder  httptransport.DecodeRequestFunc
		Encoder  httptransport.EncodeResponseFunc
	}
)
