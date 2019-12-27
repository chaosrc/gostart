package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials"
)

var _ credentials.PerRPCCredentials = (*Authentication)(nil)

type Authentication struct {
	User     string
	Password string
}

func (auth *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": auth.User, "password": auth.Password}, nil
}

func (auth *Authentication) RequireTransportSecurity() bool {
	return false
}

func (auth *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appkey string
	var appid string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != auth.User || appkey != auth.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
