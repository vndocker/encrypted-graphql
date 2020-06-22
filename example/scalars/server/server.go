package main

import (
	"log"
	"net/http"

	"github.com/vndocker/encrypted-graphql/example/scalars"
	"github.com/vndocker/encrypted-graphql/graphql/handler"
	"github.com/vndocker/encrypted-graphql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
