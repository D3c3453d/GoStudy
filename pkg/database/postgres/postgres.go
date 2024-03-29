package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Connect("pgx", connectString)
	if err != nil {
		logrus.Fatal("Connect error: ", err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatal("Ping error: ", err)
	}
	return db, nil
}
