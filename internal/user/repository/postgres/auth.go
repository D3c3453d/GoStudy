package postgres

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(account entity.Account) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, phone, description) values ($1, $2, $3, $4) RETURNING id", config.AccountsTable)

	row := r.db.QueryRow(query, account.Name, account.Password, account.Phone, account.Desc)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(name, password string) (entity.Account, error) {
	var account entity.Account
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password_hash=$2", config.AccountsTable)
	err := r.db.Get(&account, query, name, password)
	return account, err
}
