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
	// Configurar cliente de base de datos
	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}
	defer client.Close()

	// Configurar router
	router := chi.NewRouter()

	// Configurar CORS
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:           true,
	}).Handler)

	// Configurar GraphQL
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(client),
	}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Configurar servidor HTTP
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(httpServer.ListenAndServe())
}