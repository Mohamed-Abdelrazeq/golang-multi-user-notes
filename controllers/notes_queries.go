package handler

import (
	"errors"

	"github.com/Fiber-CRUD/models"
	"github.com/google/uuid"
)

func (dataSource *DataSource) queryAllNotes() (*[]models.Note, error) {
	notes := []models.Note{}

	rows, err := dataSource.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}

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

func (dataSource *DataSource) excuteInsertNote(note models.Note) error {

	_, err := dataSource.Exec("INSERT INTO notes VALUES ($1, $2, $3, $4)",
		&note.Id,
		&note.CreatedAt,
		&note.UpdatedAt,
		&note.Content,
	)

	return err

}

func (DataSource *DataSource) excuteDeleteNote(id *uuid.UUID) error {
	msg, err := DataSource.Exec("DELETE FROM notes WHERE id = $1",
		&id,
	)

	if err != nil {
		return err
	}

	count, _ := msg.RowsAffected()
	if count == 0 {
		return errors.New("note doesn't exist")
	}

	return nil
}
