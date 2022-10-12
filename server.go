package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kjais1720/graphql-go-server/graph"
	"github.com/kjais1720/graphql-go-server/graph/generated"
	"github.com/kjais1720/graphql-go-server/db"

	"github.com/go-chi/chi/v5"
)

const DEFAULT_PORT = "8080"
const DB_URI = "mongodb://127.0.0.1:27017"
const DB_NAME = "graphql"


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	router := chi.NewRouter()

	client, ctx, cancel, err := db.Connect(DB_URI)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(client, ctx, cancel)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
