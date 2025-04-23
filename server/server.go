package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"tools-review-backend/database"
	"tools-review-backend/graph"
	"tools-review-backend/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func Start() {
	// Setup data base
	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}
	defer client.Close()

	// Setup router
	router := chi.NewRouter()

	// Setup cors
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:           true,
	}).Handler)

	// Setup GraphQL
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(client),
	}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setup HTTP server
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("server running in  http://localhost:%s/ ", port)
	log.Fatal(httpServer.ListenAndServe())
}