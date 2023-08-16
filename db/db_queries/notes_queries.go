package db_queries

import (
	"github.com/Fiber-CRUD/db/db_connection"
	"github.com/Fiber-CRUD/models"
)

func GetAllNotes() (*[]models.Note, error) {
	notes := []models.Note{}

	db, err := db_connection.OpenDB()
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

	return &notes, nil
}
