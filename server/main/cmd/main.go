package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	graph "smolf-main/gql"
	"smolf-main/internal/handler/gql"
	authHandler "smolf-main/internal/handler/gql/auth"
	repoGRPC "smolf-main/internal/repository/grpc"
	userUC "smolf-main/internal/usecase/user"
	authPB "smolf-main/pb/auth"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"google.golang.org/grpc"
)

func main() {
	// Initialize grpc client
	//
	// Auth
	authGRPCClient, err := grpc.NewClient("localhost:50051", nil)
	if err != nil {
		log.Fatal(err)
	}
	authClient := authPB.NewAuthServiceClient(authGRPCClient)

	// Initialize repositories
	repositoryGRPC := repoGRPC.NewRepository(repoGRPC.RepositoryParam{
		Auth: authClient,
	})

	// Initialize usecases
	userUC := userUC.NewUsecase(repositoryGRPC)

	// Initialize handlers
	authHandlers := authHandler.NewHandler(userUC)

	// Initialize graphql server
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: gql.NewResolver(gql.ResolverParam{
				Auth: authHandlers,
			}),
		}),
	)

	httpServer := &http.Server{
		Addr: ":5002",
	}

	// Initialize graphql http handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The system is healthy")
	})
	http.Handle("/gql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	fmt.Println("HTTP Server is running at port 5002")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down the server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stopping the HTTP server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
		log.Fatal("Shutdown timeout")
	default:
		fmt.Println("Server shutted down gracefully.")
		os.Exit(0)
	}
}
