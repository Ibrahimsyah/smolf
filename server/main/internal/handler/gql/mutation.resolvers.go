package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"smolf-main/gql"
)

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, payload gql.RegisterPayload) (*gql.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Register - Register"))
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }