package internal

import "database/sql"

type Scope struct {
	db *sql.DB
}

func NewScope(DB *sql.DB) *Scope {
	s := &Scope{
		db: DB,
	}

	return s
}
