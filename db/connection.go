package db

import (
	"database/sql"
	"os"
)

type apiConfig struct {
	DB *Queries
}

var DBConnection apiConfig

func OpenDBConnection() error {
	db, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DB"))
	if err != nil {
		return err
	}
	dbQueries := New(db)

	DBConnection = apiConfig{
		DB: dbQueries,
	}

	return nil
}
