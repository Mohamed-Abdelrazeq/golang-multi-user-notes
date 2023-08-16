package db_connection

import "database/sql"

type DataSource struct {
	*sql.DB
}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
