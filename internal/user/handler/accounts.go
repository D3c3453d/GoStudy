package handler

import (
	"GoStudy/internal/config"
	"GoStudy/internal/user/entity"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Help(c *config.Commands) {
	h.services.Help(c)
}
func (h *Handler) Add() {
	var account entity.Account

	fmt.Print("Enter your username:\n")
	_, err := fmt.Scan(&account.UserName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your phone number:\n")
	_, err = fmt.Scan(&account.UserPhone)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	fmt.Print("Enter your description:\n")
	_, err = fmt.Scan(&account.UserDesc)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	err = h.services.Create(account)
	if err != nil {
		logrus.Warnln(err)
		return
	}
}

func (h *Handler) All() {
	list, err := h.services.GetAll()
	if err != nil {
		logrus.Warnln(err)
		return
	}
	for i, account := range list {
		fmt.Printf("%d:\n  Name: %s\n  Phone: %s\n  Description: %s\n",
			i+1, account.UserName, account.UserPhone, account.UserDesc)
	}
}

func (h *Handler) Phone() {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	list, err := h.services.GetByName(userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	for i, account := range list {
		fmt.Printf("%d:\n  Name: %s\n  Phone: %s\n",
			i+1, account.UserName, account.UserPhone)
	}
}

func (h *Handler) Desc() {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	list, err := h.services.GetByName(userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	if err != nil {
		logrus.Warnln(err)
		return
	}
	for i, account := range list {
		fmt.Printf("%d:\n  Name: %s\n  Description: %s\n",
			i+1, account.UserName, account.UserDesc)
	}
}

func (h *Handler) Show() {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	list, err := h.services.GetByName(userName)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	for i, account := range list {
		fmt.Printf("%d:\n  Name: %s\n  Phone: %s\n  Description: %s\n",
			i+1, account.UserName, account.UserPhone, account.UserDesc)
	}
}

func (h *Handler) Find() {
	var userPhone string
	fmt.Print("Enter phone number:\n")
	_, err := fmt.Scan(&userPhone)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	list, err := h.services.GetByPhone(userPhone)
	if err == nil {
		for _, account := range list {
			fmt.Printf("  Name: %s\n  Phone: %s\n  Description: %s\n",
				account.UserName, account.UserPhone, account.UserDesc)
		}
	} else {
		fmt.Println("Not found")
	}
}
