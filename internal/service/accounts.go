package service

import (
	"GoStudy/internal/config"
	"GoStudy/internal/model"
	"GoStudy/internal/repository/postgres"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountsService struct {
}

func Help(c *config.Commands) {
	fmt.Printf("%s to add new account\n", c.Add)
	fmt.Printf("%s to see all accounts\n", c.All)
	fmt.Printf("%s to see description of the account\n", c.Desc)
	fmt.Printf("%s to see phone number of the account\n", c.Phone)
	fmt.Printf("%s to find account by phone number\n", c.Find)
	fmt.Printf("%s to show all information about account\n", c.Show)
	fmt.Printf("%s to exit\n", c.Exit)

}

func Add(tx *sqlx.Tx) {
	var account model.Account

	fmt.Print("Enter your username:\n")
	_, err := fmt.Scan(&account.UserName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your phone number:\n")
	_, err = fmt.Scan(&account.UserPhone)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your description:\n")
	_, err = fmt.Scan(&account.UserDesc)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	postgres.Create(tx, "accounts", &account)
}

func All(db *sqlx.DB) {
	selectparam := []string{"name", "phone", "description"}
	postgres.Show(db, "accounts", selectparam)
}

func Phone(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT phone FROM accounts WHERE name=$1", userName)
	if err != nil {
		logrus.Warnln(err)
	}
	var userPhone string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userPhone)
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
	}
	err = rows.Err()

}

func Desc(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT description FROM accounts WHERE name=$1", userName)
	if err != nil {
		logrus.Warnln(err)
	}
	var userDesc string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userDesc)
		fmt.Printf("%s's description: %s\n", userName, userDesc)
	}
	err = rows.Err()
}

func Show(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT phone, description FROM accounts WHERE name=$1", userName)
	if err != nil {
		logrus.Warnln(err)
	}
	var userPhone string
	var userDesc string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userPhone, &userDesc)
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
		fmt.Printf("%s's description: %s\n", userName, userDesc)
	}
	err = rows.Err()
}

func Find(db *sqlx.DB) {
	var userPhone string
	fmt.Print("Enter phone number:\n")
	_, err := fmt.Scan(&userPhone)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	row := db.QueryRow("SELECT name FROM accounts WHERE phone=$1", userPhone)
	var userName string
	err = row.Scan(&userName)
	if err == nil {
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
	} else {
		fmt.Println("Not found")
	}
}
