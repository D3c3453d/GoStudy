package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//Commands config
type Commands struct {
	Help  string
	Add   string
	All   string
	Desc  string
	Phone string
	Find  string
	Show  string
	Exit  string
}

func LoadConfiguration(fileName string) *Commands {
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	var command Commands
	if err := viper.Unmarshal(&command); err != nil {
		fmt.Println(err)
	}
	return &command
}

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

func (d *Dict) help(c *Commands) {
	fmt.Printf("%s to add new account\n", c.Add)
	fmt.Printf("%s to see all accounts\n", c.All)
	fmt.Printf("%s to see description of the account\n", c.Desc)
	fmt.Printf("%s to see phone number of the account\n", c.Phone)
	fmt.Printf("%s to find account by phone number\n", c.Find)
	fmt.Printf("%s to show all information about account\n", c.Show)
	fmt.Printf("%s to exit\n", c.Exit)

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

	command := LoadConfiguration("./commands.json") //commands config

	dict := NewDict() //new dictionary

	//interaction
	var input string
	for {
		fmt.Scan(&input)
		switch input {
		case command.Help:
			dict.help(command)
		case command.Add:
			dict.add()
		case command.All:
			dict.all()
		case command.Phone:
			dict.phone()
		case command.Desc:
			dict.desc()
		case command.Find:
			dict.find()
		case command.Show:
			dict.show()
		case command.Exit:
			return
		default:
			fmt.Printf("%s for help\n", command.Help)
		}
	}
}
