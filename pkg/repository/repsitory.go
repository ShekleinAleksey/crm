package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Client interface {
}

type Repository struct {
	Authorization
	Client
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Client:        NewClientPostgres(db),
	}

}
