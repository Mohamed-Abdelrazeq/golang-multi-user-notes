package models

type NoteParams struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Content     string `json:"content" validate:"required,min=1,max=65535"`
	ImageUrl    string `json:"image_url"`
	IsFavourite bool   `json:"is_favourite"`
}
