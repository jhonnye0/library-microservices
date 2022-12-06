package endpoints

import (
	"context"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"

	"github.com/jhonnye0/backend-library/pkg/authorization"
)

type Set struct {
	LoginEndpoint  endpoint.Endpoint
	LogoutEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc authorization.Service) Set {
	return Set{
		LoginEndpoint:  MakeAuthEndpoint(svc),
		LogoutEndpoint: MakeLogoutEndpoint(svc),
	}
}

func MakeAuthEndpoint(svc authorization.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := svc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return LoginResponse{token, err.Error()}, nil
		}
		return LoginResponse{token, ""}, nil
	}
}

func MakeLogoutEndpoint(svc authorization.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogoutRequest)
		err := svc.Logout(ctx, req.Token)
		if err != nil {
			return LoginResponse{"couldn't logout from th app", err.Error()}, nil
		}
		return LoginResponse{"logout successfully!", ""}, nil
	}
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
