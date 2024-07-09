package grpc

import "smolf-auth/internal/handler/grpc/auth"

type HandlerWrapper struct {
	Auth *auth.Handler
}

func NewHandlerWrapper(auth *auth.Handler) *HandlerWrapper {
	return &HandlerWrapper{
		Auth: auth,
	}
}
