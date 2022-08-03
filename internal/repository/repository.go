package repository

import (
	"GoStudy/internal/model"
	"github.com/jmoiron/sqlx"
)

type Accounts interface {
	Create(tx *sqlx.Tx, account *model.Account)
	Show(db *sqlx.DB, selectparam string, whereparam string)
}
