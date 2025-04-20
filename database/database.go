package database

import (
	"context"
	"tools-review-backend/ent"

	_ "github.com/lib/pq"
)

const (
	connStr = "host=localhost port=5432 user=admin dbname=tools-back password=1234 sslmode=disable"
)

// NewClient crea un nuevo cliente de Ent
func NewClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ejecutar migraciones
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}