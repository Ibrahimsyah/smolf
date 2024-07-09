package app

import (
	"smolf-auth/internal/handler/grpc"
	"smolf-auth/internal/handler/grpc/auth"
	"smolf-auth/internal/handler/http"
)

type Handler struct {
	GRPC *grpc.HandlerWrapper
	HTTP *http.Handler
}

func NewHandler(usecases *Usecases) *Handler {
	return &Handler{
		GRPC: grpc.NewHandlerWrapper(
			auth.NewHandler(usecases.User),
		),
		HTTP: http.NewHandler(),
	}
}
