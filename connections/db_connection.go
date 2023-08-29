package connections

import (
	"database/sql"
	"os"

	"github.com/multi-user-notes-app/db"
)

type apiConfig struct {
	DB *db.Queries
}

var DBConnection apiConfig

func OpenDBConnection() error {
	sqlDB, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DB"))
	if err != nil {
		return err
	}
	dbQueries := db.New(sqlDB)

	DBConnection = apiConfig{
		DB: dbQueries,
	}

	return nil
}
