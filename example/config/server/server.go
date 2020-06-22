package main

import (
	"log"
	"net/http"

	todo "github.com/vndocker/encrypted-graphql/example/config"
	"github.com/vndocker/encrypted-graphql/graphql/handler"
	"github.com/vndocker/encrypted-graphql/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
