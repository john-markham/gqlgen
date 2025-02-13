package main

import (
	"log"
	"net/http"

	"github.com/john-markham/gqlgen/_examples/dataloader"
	"github.com/john-markham/gqlgen/graphql/handler"
	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/john-markham/gqlgen/graphql/playground"
)

func main() {
	router := http.NewServeMux()

	srv := handler.New(
		dataloader.NewExecutableSchema(dataloader.Config{Resolvers: &dataloader.Resolver{}}),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	router.Handle("/", playground.Handler("Dataloader", "/query"))
	router.Handle("/query", srv)

	log.Println("connect to http://localhost:8082/ for graphql playground")
	log.Fatal(http.ListenAndServe(":8082", dataloader.LoaderMiddleware(router)))
}
