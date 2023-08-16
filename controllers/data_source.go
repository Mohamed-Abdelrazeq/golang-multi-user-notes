package handler

import "database/sql"

type DataSource struct {
	*sql.DB
}
