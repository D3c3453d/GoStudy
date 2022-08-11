package server

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/handler"
	"GoStudy/internal/user/repository"
	"GoStudy/internal/user/service"
	"GoStudy/pkg/database/postgres"
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewApp(config *config.Config) {
	db := initPostgresDB(&config.DBConfig)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatal("Server run error ", err)
	}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func initPostgresDB(config *config.DBConfig) *sqlx.DB {
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     config.Host,
		Port:     config.Port,
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
