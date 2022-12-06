package endpoints

import (
	"context"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"

	"github.com/jhonnye0/backend-library/pkg/hello"
)

type Set struct {
	HelloEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc hello.Service) Set {
	return Set{
		HelloEndpoint: MakeHelloEndpoint(svc),
	}
}

func MakeHelloEndpoint(svc hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		phrase := svc.Hello(ctx, req.Name)
		return HelloResponse{phrase, ""}, nil
	}
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
