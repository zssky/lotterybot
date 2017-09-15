package db

import (
	"database/sql"
	"github.com/juju/errors"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3 struct {
	*sql.DB
}

func NewSqlite3(file string) (*Sqlite3, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &Sqlite3{db}, nil
}

func (s *Sqlite3) Close() {
	s.Close()
}
