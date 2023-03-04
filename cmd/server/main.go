package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/discentral/backend/internal/query"
	"github.com/discentral/backend/pkg/config"
	"github.com/discentral/backend/pkg/graphiql"
	"github.com/discentral/backend/pkg/sdl"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/joho/godotenv/autoload"
)

var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(sdl.Schema, &query.Resolver{})
}

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.GetWithFallback("PRODUCTION", "false") == "true" {
			fmt.Fprint(w, "Hello World!")
		} else {
			w.Write(graphiql.Page)
		}
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
