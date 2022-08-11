package service

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository"
)

type AuthServiceI interface {
	Create(account entity.Account) (int, error)
	GenerateToken(name string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type AccountsServiceI interface {
	GetAll() ([]entity.Account, error)
	GetByName(name string) ([]entity.Account, error)
	GetByPhone(phone string) ([]entity.Account, error)
}

type Service struct {
	AccountsServiceI
	AuthServiceI
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AccountsServiceI: NewAccountsService(repos.AccountsReposI),
		AuthServiceI:     NewAuthService(repos.AuthReposI),
	}
}
