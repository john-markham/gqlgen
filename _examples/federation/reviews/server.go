//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/_examples/federation/reviews/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
