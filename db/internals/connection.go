package internals

import (
	"database/sql"
	"log"
	"os"
)

type apiConfig struct {
	DB *Queries
}

var DBConnection apiConfig

func InitDB() {
	db, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DB"))
	if err != nil {
		log.Fatal(err.Error())
	}
	dbQueries := New(db)

	DBConnection = apiConfig{
		DB: dbQueries,
	}
}
