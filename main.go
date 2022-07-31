package main

import (
	"GoStudy/pkg/repository"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func help(c *Commands) {
	fmt.Printf("%s to add new account\n", c.Add)
	fmt.Printf("%s to see all accounts\n", c.All)
	fmt.Printf("%s to see description of the account\n", c.Desc)
	fmt.Printf("%s to see phone number of the account\n", c.Phone)
	fmt.Printf("%s to find account by phone number\n", c.Find)
	fmt.Printf("%s to show all information about account\n", c.Show)
	fmt.Printf("%s to exit\n", c.Exit)

}

func add(tx *sqlx.Tx) {
	var account Account

	fmt.Print("Enter your username:\n")
	_, err := fmt.Scan(&account.userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your phone number:\n")
	_, err = fmt.Scan(&account.userPhone)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your description:\n")
	_, err = fmt.Scan(&account.userDesc)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	tx.Exec("INSERT INTO accounts (name, phone, description) VALUES ($1, $2, $3)", account.userName, account.userPhone, account.userDesc)
	err = tx.Commit()
	if err != nil {
		logrus.Warnln(err)
	}
}

func all(db *sqlx.DB) {
	rows, err := db.Query("SELECT name FROM accounts")
	if err != nil {
		logrus.Warnln(err)
	}
	var name string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Println(name)
	}
	err = rows.Err()
}

func phone(db *sqlx.DB) {
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

func desc(db *sqlx.DB) {
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

func show(db *sqlx.DB) {
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

func find(db *sqlx.DB) {
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

func main() {
	command := NewCommandsConf("./commands.env")
	dbconf := NewDBConf("./db.env")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "172.21.0.2",
		Port:     "5432",
		Username: dbconf.Username,
		Password: dbconf.Password,
		DBName:   dbconf.DBName,
		SSLMode:  "disable",
	})
	tx := db.MustBegin()
	if err != nil {
		logrus.Fatal("Cant create", err)
	}

	//interaction
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			logrus.Warnln(err)
		}
		switch input {
		case command.Help:
			help(command)
		case command.Add:
			add(tx)
		case command.All:
			all(db)
		case command.Phone:
			phone(db)
		case command.Desc:
			desc(db)
		case command.Find:
			find(db)
		case command.Show:
			show(db)
		case command.Exit:
			return
		default:
			fmt.Printf("%s for help\n", command.Help)
		}
	}
}
