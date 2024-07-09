package gql

import "smolf-main/internal/handler/gql/auth"

type Resolver struct {
	Auth *auth.Handler
}

type ResolverParam struct {
	Auth *auth.Handler
}

func NewResolver(param ResolverParam) *Resolver {
	return &Resolver{
		Auth: param.Auth,
	}
}
