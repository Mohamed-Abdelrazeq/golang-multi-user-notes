package core

import "database/sql"

type DataSource struct {
	*sql.DB
}
