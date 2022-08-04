package service

import (
	"GoStudy/internal/config"
	"GoStudy/internal/model"
	"GoStudy/internal/repository"
)

type Accounts interface {
	Help(c *config.Commands)
	Create(account model.Account) error
	GetAll() ([]model.Account, error)
	GetByName(name string) ([]model.Account, error)
	GetByPhone(phone string) ([]model.Account, error)
}

type Service struct {
	Accounts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Accounts: NewAccountsService(repos.Accounts),
	}
}
