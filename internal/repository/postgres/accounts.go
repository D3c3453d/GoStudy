package postgres

import (
	"GoStudy/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

func Create(tx *sqlx.Tx, table string, account *model.Account) {
	querystring := fmt.Sprintf(
		"INSERT INTO %s (name, phone, description) VALUES (%s, %s, %s)",
		table, account.UserName, account.UserPhone, account.UserDesc)
	_, err := tx.Exec(querystring)
	if err != nil {
		logrus.Warnln(err)
	}
	err = tx.Commit()
	if err != nil {
		logrus.Warnln(err)
	}
}

func Show(db *sqlx.DB, table string, selectparam []string) {
	querystring := fmt.Sprintf(
		"SELECT %s FROM %s",
		strings.Join(selectparam, ", "), table)
	rows, err := db.Query(querystring)
	if err != nil {
		logrus.Warnln(err)
	}
	data := make([]string, len(selectparam))
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&data)
		fmt.Printf("%s\n", strings.Join(data, " "))
	}
	err = rows.Err()
}
