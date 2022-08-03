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
			service.Help(&command)
		case command.Add:
			service.Add(tx)
		case command.All:
			service.All(db)
		case command.Phone:
			service.Phone(db)
		case command.Desc:
			service.Desc(db)
		case command.Find:
			service.Find(db)
		case command.Show:
			service.Show(db)
		case command.Exit:
			return
		default:
			fmt.Printf("%s for help\n", command.Help)
		}
	}
}
