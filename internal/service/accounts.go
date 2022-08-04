package service

import (
	"GoStudy/internal/config"
	"GoStudy/internal/model"
	"GoStudy/internal/repository"
	"fmt"
)

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	return &AccountsService{repo: repo}
}

func (s *AccountsService) Help(c *config.Commands) {
	fmt.Printf("%s to add new account\n", c.Add)
	fmt.Printf("%s to see all accounts\n", c.All)
	fmt.Printf("%s to see description of the account\n", c.Desc)
	fmt.Printf("%s to see phone number of the account\n", c.Phone)
	fmt.Printf("%s to find account by phone number\n", c.Find)
	fmt.Printf("%s to show all information about account\n", c.Show)
	fmt.Printf("%s to exit\n", c.Exit)

}

func (s *AccountsService) Create(account model.Account) error {
	return s.repo.Create(account)
}

func (s *AccountsService) GetAll() ([]model.Account, error) {
	return s.repo.GetAll()
}

func (s *AccountsService) GetByName(name string) ([]model.Account, error) {
	return s.repo.GetByName(name)
}

func (s *AccountsService) GetByPhone(phone string) ([]model.Account, error) {
	return s.repo.GetByPhone(phone)
}
