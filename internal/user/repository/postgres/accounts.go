package postgres

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountsPostgres struct {
	db *sqlx.DB
}

func NewAccountsPostgres(db *sqlx.DB) *AccountsPostgres {
	return &AccountsPostgres{db: db}
}

func (r *AccountsPostgres) GetAll() ([]entity.Account, error) {
	var list []entity.Account

	query := fmt.Sprintf("SELECT * FROM %s", config.AccountsTable)
	err := r.db.Select(&list, query)

	return list, err
}

func (r *AccountsPostgres) GetByName(name string) ([]entity.Account, error) {
	var list []entity.Account

	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", config.AccountsTable)
	err := r.db.Select(&list, query, name)

	return list, err
}

func (r *AccountsPostgres) GetByPhone(phone string) ([]entity.Account, error) {
	var list []entity.Account

	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=$1", config.AccountsTable)
	err := r.db.Select(&list, query, phone)

	return list, err
}
