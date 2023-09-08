package models

type CreateNoteParams struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Content     string `json:"content" validate:"required,min=1,max=65535"`
	ImageUrl    string `json:"image_url"`
	IsFavourite bool   `json:"is_favourite"`
}

type NoteDetailsParams struct {
	ID int32 `json:"id" validate:"required"`
}

type UpdateNoteParams struct {
	NoteDetailsParams
	CreateNoteParams
}

type UpdateFavouriteParams struct {
	NoteDetailsParams
}
