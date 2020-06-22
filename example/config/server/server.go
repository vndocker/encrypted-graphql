package main

import (
	"log"
	"net/http"

	todo "github.com/LaBanHSPO/gqlgen/example/config"
	"github.com/LaBanHSPO/gqlgen/graphql/handler"
	"github.com/LaBanHSPO/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
