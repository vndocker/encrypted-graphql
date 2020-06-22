package main

import (
	"log"
	"net/http"

	"github.com/LaBanHSPO/gqlgen/example/scalars"
	"github.com/LaBanHSPO/gqlgen/graphql/handler"
	"github.com/LaBanHSPO/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
