package main

import "fmt"

type Account struct {
	userName  string
	userPhone string
	userDesc  string
}

type Dict struct {
	dict map[string]Account
}

func NewDict() *Dict {
	var d Dict
	d.dict = make(map[string]Account)
	return &d
}

func (d *Dict) help() {
	fmt.Println("/add to add new account")
	fmt.Println("/all to see all accounts")
	fmt.Println("/desc to see description of the account")
	fmt.Println("/phone to see phone number of the account")
	fmt.Println("/find to find by phone number")
}

func (d *Dict) add() {
	var account Account

	fmt.Print("Enter your username:\n")
	fmt.Scan(&account.userName)
	fmt.Print("Enter your phone number:\n")
	fmt.Scan(&account.userPhone)
	fmt.Print("Enter your description:\n")
	fmt.Scan(&account.userDesc)
	d.dict[account.userName] = account
}

func (d *Dict) all() {
	for userName := range d.dict {
		fmt.Println(userName)
	}
}

func (d *Dict) phone() {
	var userName string
	fmt.Print("Enter username:\n")
	fmt.Scan(&userName)
	fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
}

func (d *Dict) desc() {
	var userName string
	fmt.Print("Enter username:\n")
	fmt.Scan(&userName)
	fmt.Printf("%s's description: %s\n", userName, d.dict[userName].userDesc)
}

func (d *Dict) show() {
	var userName string

	fmt.Print("Enter username:\n")
	fmt.Scan(&userName)
	fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
	fmt.Printf("%s's description: %s\n", userName, d.dict[userName].userDesc)
}

func (d *Dict) find() {
	var userPhone string
	fmt.Print("Enter phone number:\n")
	fmt.Scan(&userPhone)
	for userName := range d.dict {
		if userPhone == d.dict[userName].userPhone {
			fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
			return
		}
	}
	fmt.Println("Not found")
}

func main() {
	dict := NewDict()
	var input string
	for {
		fmt.Scan(&input)
		switch input {
		case "/help":
			dict.help()
		case "/add":
			dict.add()
		case "/all":
			dict.all()
		case "/phone":
			dict.phone()
		case "/desc":
			dict.desc()
		case "/find":
			dict.find()
		case "/show":
			dict.show()
		case "/exit":
			break
		default:
			fmt.Println("/help for help")
		}
	}
}
