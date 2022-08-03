package main

import (
	"GoStudy/internal/config"
	"GoStudy/internal/service"
	"GoStudy/pkg/database/postgres"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

func main() {
	var command config.Commands
	var dbconf config.DBConfig
	command.LoadConfig("./commands.env")
	dbconf.LoadConfig("./db.env")

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "db",
		Port:     "5432",
		Username: dbconf.Username,
		Password: dbconf.Password,
		DBName:   dbconf.DBName,
		SSLMode:  "disable",
	})
	tx := db.MustBegin()
	if err != nil {
		logrus.Fatal("Cant create ", err)
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
