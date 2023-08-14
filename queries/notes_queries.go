package queries

import (
	"database/sql"

	"github.com/Fiber-CRUD/models"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:5024@localhost:5432/Fiber-CRUD?sslmode=disable",
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllNotes() ([]models.Note, error) {
	notes := []models.Note{}

	db, err := openDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		note := new(models.Note)
		rows.Scan(
			&note.Id,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.Content,
		)
		notes = append(notes, *note)
	}

	return notes, nil
}
