package auth

import (
	"context"
	"smolf-auth/pb/auth"
)

func (h *Handler) AuthByEmailAndPassword(ctx context.Context, in *auth.AuthEmailPasswordRequest) (*auth.AuthResponse, error) {
	return &auth.AuthResponse{
		Token:  "token",
		UserId: 1241,
	}, nil
}
