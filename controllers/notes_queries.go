package handler

import (
	"github.com/Fiber-CRUD/models"
)

func (dataSource *DataSource) QueryAllNotes() (*[]models.Note, error) {
	notes := []models.Note{}

	rows, err := dataSource.Query("SELECT * FROM notes")
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
