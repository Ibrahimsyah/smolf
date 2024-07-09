package auth

import (
	"smolf-auth/internal/usecase/user"
	"smolf-auth/pb/auth"
)

type Handler struct {
	UserUC *user.Usecase

	auth.UnimplementedAuthServiceServer
}

func NewHandler(userUC *user.Usecase) *Handler {
	return &Handler{
		UserUC: userUC,
	}
}
