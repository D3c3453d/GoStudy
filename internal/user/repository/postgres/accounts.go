package postgres

import (
	"GoStudy/internal/user/entity"
	"GoStudy/pkg/database/postgres"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountsPostgres struct {
	db *sqlx.DB
}

func NewAccountsPostgres(db *sqlx.DB) *AccountsPostgres {
	return &AccountsPostgres{db: db}
}

func (r *AccountsPostgres) Create(account model.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createAccountsQuery := fmt.Sprintf(
		"INSERT INTO %s (name, phone, description) VALUES ($1, $2, $3) RETURNING id", postgres.AccountsTable)
	row := tx.QueryRow(createAccountsQuery, account.UserName, account.UserPhone, account.UserDesc)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *AccountsPostgres) GetAll() ([]model.Account, error) {
	var list []model.Account

	query := fmt.Sprintf("SELECT * FROM %s", postgres.AccountsTable)
	err := r.db.Select(&list, query)

	return list, err
}

func (r *AccountsPostgres) GetByName(name string) ([]model.Account, error) {
	var list []model.Account

	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", postgres.AccountsTable)
	err := r.db.Select(&list, query, name)

	return list, err
}

func (r *AccountsPostgres) GetByPhone(phone string) ([]model.Account, error) {
	var list []model.Account

	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=$1", postgres.AccountsTable)
	err := r.db.Select(&list, query, phone)

	return list, err
}
