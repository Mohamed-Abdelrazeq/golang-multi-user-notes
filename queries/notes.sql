-- CreateNote: one
INSERT INTO notes (content) VALUES ($1) RETURNING *;