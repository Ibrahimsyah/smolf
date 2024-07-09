package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"smolf-auth/internal/app"
	"smolf-auth/pb/auth"
	"syscall"
	"time"

	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize external dependencies

	// Initialize layers
	repositories := app.NewRepositories(app.RepositoryParam{})
	usecases := app.NewUsecase(repositories)
	handlers := app.NewHandler(usecases)

	server := gogrpc.NewServer()

	reflection.Register(server)
	auth.RegisterAuthServiceServer(server, handlers.GRPC.Auth)

	net, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Starting GRPC server
	go func() {
		if err := server.Serve(net); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("GRPC Server is running at port 50051")

	// Starting HTTP server
	httpServer := &http.Server{
		Addr:    ":5001",
		Handler: http.DefaultServeMux,
	}
	handlers.HTTP.RegisterRouters()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	fmt.Println("HTTP Server is running at port 5001")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down the server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stopping the GRPC server
	server.GracefulStop()

	// Stopping the HTTP server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		log.Fatal("Shutdown timeout")
	default:
		fmt.Println("Server shutted down gracefully.")
	}
}
