package main

import (
	"log"
	"net/http"
	"os"

	"github.com/landrade/gqlgen-cache-control-plugin/_example/graph"
	"github.com/landrade/gqlgen-cache-control-plugin/_example/graph/generated"
	"github.com/landrade/gqlgen-cache-control-plugin/cache"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func new() *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(cache.Extension{})
	return srv
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := new()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cache.Middleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
