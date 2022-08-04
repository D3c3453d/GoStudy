package repository

import (
	"GoStudy/internal/model"
	"GoStudy/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Accounts interface {
	Create(account model.Account) error
	GetAll() ([]model.Account, error)
	GetByName(name string) ([]model.Account, error)
	GetByPhone(phone string) ([]model.Account, error)
}

type Repository struct {
	Accounts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Accounts: postgres.NewAccountsPostgres(db),
	}
}
