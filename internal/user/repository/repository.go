package repository

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type AuthReposI interface {
	Create(account entity.Account) (int, error)
	GetUser(username, password string) (entity.Account, error)
}

type AccountsReposI interface {
	GetAll() ([]entity.Account, error)
	GetByName(name string) ([]entity.Account, error)
	GetByPhone(phone string) ([]entity.Account, error)
}

type Repository struct {
	AccountsReposI
	AuthReposI
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountsReposI: postgres.NewAccountsPostgres(db),
		AuthReposI:     postgres.NewAuthPostgres(db),
	}
}
