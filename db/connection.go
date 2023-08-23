package db

import (
	"database/sql"
)

type apiConfig struct {
	DB *Queries
}

var DBConnection apiConfig

func OpenDBConnection() error {
	db, err := sql.Open("postgres", "postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable")
	if err != nil {
		return err
	}
	dbQueries := New(db)

	DBConnection = apiConfig{
		DB: dbQueries,
	}

	return nil
}
