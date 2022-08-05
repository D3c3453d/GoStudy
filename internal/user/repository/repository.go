package repository

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type AccountsReposI interface {
	Create(account entity.Account) error
	GetAll() ([]entity.Account, error)
	GetByName(name string) ([]entity.Account, error)
	GetByPhone(phone string) ([]entity.Account, error)
}

type Repository struct {
	AccountsReposI
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountsReposI: postgres.NewAccountsPostgres(db),
	}
}
