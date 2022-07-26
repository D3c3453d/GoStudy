package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Commands config
type Commands struct {
	Help  string `mapstructure:"HELP"`
	Add   string `mapstructure:"ADD"`
	All   string `mapstructure:"ALL"`
	Desc  string `mapstructure:"DESC"`
	Phone string `mapstructure:"PHONE"`
	Find  string `mapstructure:"FIND"`
	Show  string `mapstructure:"SHOW"`
	Exit  string `mapstructure:"EXIT"`
}

func LoadConfiguration(fileName string) *Commands {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Read file error", err)
	}
	var command Commands
	if err := viper.Unmarshal(&command); err != nil {
		log.Panic("Parse file error", err)
	}
	return &command
}

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=db port=5432 user=postgresuser dbname=postgresdb password=qwerty sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

type Account struct {
	userName  string `db:"name""`
	userPhone string `db:"phone""`
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

func add(tx *sqlx.Tx) {
	var account Account

	fmt.Print("Enter your username:\n")
	_, err := fmt.Scan(&account.userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	fmt.Print("Enter your phone number:\n")
	_, err = fmt.Scan(&account.userPhone)
	if err != nil {
		log.Warnln(err)
		return
	}
	fmt.Print("Enter your description:\n")
	_, err = fmt.Scan(&account.userDesc)
	if err != nil {
		log.Warnln(err)
		return
	}
	tx.MustExec("INSERT INTO accounts (name, phone, description) VALUES ($1, $2, $3)", "account.userName", "account.userPhone", "account.userDesc")
	err = tx.Commit()
	if err != nil {
		log.Warnln(err)
	}
}

func (d *Dict) all() {
	for userName := range d.dict {
		fmt.Println(userName)
	}
}

func (d *Dict) phone() {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
}

func (d *Dict) desc() {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	fmt.Printf("%s's description: %s\n", userName, d.dict[userName].userDesc)
}

func (d *Dict) show() {
	var userName string

	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
	fmt.Printf("%s's description: %s\n", userName, d.dict[userName].userDesc)
}

func (d *Dict) find() {
	var userPhone string
	fmt.Print("Enter phone number:\n")
	_, err := fmt.Scan(&userPhone)
	if err != nil {
		log.Warnln(err)
		return
	}
	for userName := range d.dict {
		if userPhone == d.dict[userName].userPhone {
			fmt.Printf("%s's phone number: %s\n", userName, d.dict[userName].userPhone)
			return
		}
	}
	fmt.Println("Not found")
}

func main() {
	db, err := NewPostgresDB()
	tx := db.MustBegin()
	if err != nil {
		log.Fatal(err)
	}

	command := LoadConfiguration("./commands.env") //commands config

	dict := NewDict() //new dictionary

	//interaction
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Warnln(err)
		}
		switch input {
		case command.Help:
			dict.help(command)
		case command.Add:
			add(tx)
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
