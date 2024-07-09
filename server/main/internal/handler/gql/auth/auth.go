package auth

import (
	"context"
	"smolf-main/gql"
)

func (h *Handler) HandleLogin(ctx context.Context, payload gql.LoginPayload) (*gql.AuthResponse, error) {
	return &gql.AuthResponse{
		UserID: 11223,
		Token:  "token",
	}, nil
}
