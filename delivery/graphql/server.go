package graphql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/stdapps/graphql-example/delivery/graphql/graph/generated"
	"github.com/stdapps/graphql-example/delivery/graphql/graph/resolvers"
	"github.com/stdapps/graphql-example/storage"
)

// StartGraphqlServer starts a graphql server and blocks
func StartGraphqlServer(port string, db storage.PostgresStorage) {
	resolver := &resolvers.Resolver{
		Storage: db,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
