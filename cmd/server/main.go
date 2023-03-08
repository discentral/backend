package main

import (
	"log"
	"net/http"

	"github.com/discentral/backend/internal/root"
	"github.com/discentral/backend/pkg/config"
	"github.com/discentral/backend/pkg/graphiql"
	"github.com/discentral/backend/pkg/sdl"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/joho/godotenv/autoload"
)

var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(sdl.Schema, &root.Resolver{})
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.GetWithFallback("PRODUCTION", "true") == "true" {
			w.Write([]byte("Hello World!"))
		} else {
			w.Write(graphiql.Page)
		}
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
