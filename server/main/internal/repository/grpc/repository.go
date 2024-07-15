package grpc

import (
	"context"

	pbAuth "smolf-main/pb/auth"

	"google.golang.org/grpc"
)

type AuthGRPCProvider interface {
	AuthByEmailAndPassword(ctx context.Context, in *pbAuth.AuthEmailPasswordRequest, opts ...grpc.CallOption) (*pbAuth.AuthResponse, error)
}

type Repository struct {
	Auth AuthGRPCProvider
}

type RepositoryParam struct {
	Auth AuthGRPCProvider
}

func NewRepository(param RepositoryParam) *Repository {
	return &Repository{
		Auth: param.Auth,
	}
}
