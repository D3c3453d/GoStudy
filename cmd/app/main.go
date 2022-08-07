package main

import (
	"GoStudy/internal/config"
	"GoStudy/server"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := config.NewConfig()
	server.ConsoleApp(cfg)
}
