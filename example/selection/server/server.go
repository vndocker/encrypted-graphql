package main

import (
	"log"
	"net/http"

	"github.com/vndocker/encrypted-graphql/example/selection"
	"github.com/vndocker/encrypted-graphql/graphql/handler"
	"github.com/vndocker/encrypted-graphql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
