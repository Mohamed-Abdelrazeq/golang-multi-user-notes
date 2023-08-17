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

func (dataSource *DataSource) executeInsertNote(note models.Note) error {

	_, err := dataSource.Exec("INSERT INTO notes VALUES ($1, $2, $3, $4)",
		&note.Id,
		&note.CreatedAt,
		&note.UpdatedAt,
		&note.Content,
	)

	return err
}

func (dataSource *DataSource) executeDeleteNote(id *uuid.UUID) error {
	msg, err := dataSource.Exec("DELETE FROM notes WHERE id = $1",
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

func (dataSource *DataSource) queryGetNoteById(id *uuid.UUID) (*models.Note, error) {
	rows, err := dataSource.Query("SELECT * FROM notes WHERE id = $1 LIMIT 1",
		&id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("note doesn't exist")
	}

	note := new(models.Note)
	rows.Scan(
		&note.Id,
		&note.CreatedAt,
		&note.UpdatedAt,
		&note.Content,
	)

	return note, nil
}

func (dataSource *DataSource) executeUpdateNote(note models.Note) error {
	msg, err := dataSource.Exec("UPDATE notes SET content=$2, updated_at=$3 WHERE id=$1",
		&note.Id,
		&note.Content,
		&note.UpdatedAt,
	)

	count, _ := msg.RowsAffected()
	if count == 0 {
		return errors.New("note doesn't exist")
	}

	return err
}
