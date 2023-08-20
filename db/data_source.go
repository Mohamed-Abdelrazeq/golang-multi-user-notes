package db

import "database/sql"

type DataSource struct {
	*sql.DB
}

var DBConnection DataSource

func OpenDB() error {
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable",
	)
	if err != nil {
		return err
	}

	DBConnection = DataSource{conn}

	return nil
}
