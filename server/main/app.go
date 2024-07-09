package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"smolf-main/pb/auth"

	"google.golang.org/grpc"
)

const port = 50051

type authHandler struct {
	auth.UnimplementedAuthServiceServer
}

func (h *authHandler) AuthByEmailAndPassword(ctx context.Context, param *auth.AuthEmailPasswordRequest) (*auth.AuthResponse, error) {
	return &auth.AuthResponse{}, nil
}

func main() {
	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, &authHandler{})

	http, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("failed to initialize server", err)
	}

	fmt.Println("Server is running")
	if err := server.Serve(http); err != nil {
		log.Fatal("failed to serve server", err)
	}
}
