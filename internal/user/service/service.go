package service

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository"
)

type AccountsServiceI interface {
	Help(c *config.Commands)
	Create(account entity.Account) error
	GetAll() ([]entity.Account, error)
	GetByName(name string) ([]entity.Account, error)
	GetByPhone(phone string) ([]entity.Account, error)
}

type Service struct {
	AccountsServiceI
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AccountsServiceI: NewAccountsService(repos.AccountsReposI),
	}
}
