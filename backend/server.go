
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"backend/graph"
	"backend/graph/generated"
)

const defaultPort = "8080"

func main() {
	db := graph.InitDB()
	resolver := &graph.Resolver{DB: db}


	srv := handler.New(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://wallet-log.vercel.app"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	http.Handle("/", c.Handler(playground.Handler("GraphQL playground", "/query")))

	http.Handle("/query", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "failed to read request body", http.StatusBadRequest)
				return
			}
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			var op struct{ Query string `json:"query"` }
			_ = json.Unmarshal(bodyBytes, &op)

			if strings.Contains(op.Query, "register") || strings.Contains(op.Query, "login") {
				srv.ServeHTTP(w, r)
				return
			}
		}
		graph.AuthMiddleware(srv).ServeHTTP(w, r)
	})))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
