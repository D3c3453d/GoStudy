package server

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/handler"
	"GoStudy/internal/user/repository"
	"GoStudy/internal/user/service"
	"GoStudy/pkg/database/postgres"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func ConsoleApp(config *config.Config) {
	db := initPostgresDB(&config.DBConfig)
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
		case config.Commands.Help:
			handlers.Help(&config.Commands)
		case config.Commands.Add:
			handlers.Add()
		case config.Commands.All:
			handlers.All()
		case config.Commands.Phone:
			handlers.Phone()
		case config.Commands.Desc:
			handlers.Desc()
		case config.Commands.Find:
			handlers.Find()
		case config.Commands.Show:
			handlers.Show()
		case config.Commands.Exit:
			return
		default:
			fmt.Printf("%s for help\n", config.Commands.Help)
		}
	}
}

func initPostgresDB(config *config.DBConfig) *sqlx.DB {
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     "db",
		Port:     "5432",
		Username: config.Username,
		Password: config.Password,
		DBName:   config.DBName,
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatal("Cant connect ", err)
	}
	return db
}
