package main

import (
	"GoStudy/internal/config"
	"GoStudy/internal/server"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := config.NewConfig()
	server.NewApp(cfg)
}
