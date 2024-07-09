package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	graph "smolf-main/gql"
	"smolf-main/internal/handler/gql"
	"smolf-main/internal/handler/gql/auth"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: gql.NewResolver(gql.ResolverParam{
				Auth: &auth.Handler{},
			}),
		}),
	)

	httpServer := &http.Server{
		Addr: ":5002",
	}

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
