package service

import (
	"GoStudy/internal/user/entity"
	"GoStudy/internal/user/repository"
)

type AccountsService struct {
	repo repository.AccountsReposI
}

func NewAccountsService(repo repository.AccountsReposI) *AccountsService {
	return &AccountsService{repo: repo}
}

func (s *AccountsService) GetAll() ([]entity.Account, error) {
	return s.repo.GetAll()
}

func (s *AccountsService) GetByName(name string) ([]entity.Account, error) {
	return s.repo.GetByName(name)
}

func (s *AccountsService) GetByPhone(phone string) ([]entity.Account, error) {
	return s.repo.GetByPhone(phone)
}
