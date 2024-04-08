package repository

import "database/sql"

type Repository interface {
}

type repositoryPostgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repositoryPostgres {
	return &repositoryPostgres{
		db: db,
	}
}
