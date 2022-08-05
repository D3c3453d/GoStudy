package main

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/handler"
	"GoStudy/internal/user/repository"
	"GoStudy/internal/user/service"
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
	if err != nil {
		logrus.Fatal("Cant create ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//interaction
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			logrus.Warnln(err)
		}
		switch input {
		case command.Help:
			handlers.Help(&command)
		case command.Add:
			handlers.Add()
		case command.All:
			handlers.All()
		case command.Phone:
			handlers.Phone()
		case command.Desc:
			handlers.Desc()
		case command.Find:
			handlers.Find()
		case command.Show:
			handlers.Show()
		case command.Exit:
			return
		default:
			fmt.Printf("%s for help\n", command.Help)
		}
	}
}
