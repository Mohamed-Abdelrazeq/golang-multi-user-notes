package main

import "database/sql"

type DataSource struct {
	*sql.DB
}
