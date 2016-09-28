package main

import (
    "net/http"
    "github.com/graphql-go/handler"
    "github.com/object88/relay/examples/starwars"
)

// taken from: https://github.com/graphql-go/playground
// simplest relay-compliant graphql server HTTP handler
// using Starwars schema from `graphql-relay-go` examples
func main() {
    h := handler.New(&handler.Config{
        Schema: &starwars.Schema,
        Pretty: true,
    })

    http.Handle("/graphql", h)
    http.ListenAndServe(":3000", nil)
}
