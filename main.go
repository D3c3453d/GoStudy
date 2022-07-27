package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
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
	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s target_session_attrs=read-write",
		"172.28.0.2", 5432, "postdb", "postuser", "qwerty")
	db, err := sqlx.Connect("pgx", connstring)
	if err != nil {
		log.Fatal("Connect error: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error: ", err)
	}

	return db, nil
}

type Account struct {
	userName  string `db:"name"`
	userPhone string `db:"phone"`
	userDesc  string `db:"description"`
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
	tx.Exec("INSERT INTO accounts (name, phone, description) VALUES ($1, $2, $3)", account.userName, account.userPhone, account.userDesc)
	err = tx.Commit()
	if err != nil {
		log.Warnln(err)
	}
}

func (d *Dict) all(db *sqlx.DB) {
	rows, err := db.Query("SELECT name FROM accounts")
	if err != nil {
		log.Warnln(err)
	}
	var name string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Println(name)
	}
	err = rows.Err()
}

func (d *Dict) phone(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT phone FROM accounts WHERE name=$1", userName)
	if err != nil {
		log.Warnln(err)
	}
	var userPhone string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userPhone)
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
	}
	err = rows.Err()

}

func (d *Dict) desc(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT description FROM accounts WHERE name=$1", userName)
	if err != nil {
		log.Warnln(err)
	}
	var userDesc string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userDesc)
		fmt.Printf("%s's description: %s\n", userName, userDesc)
	}
	err = rows.Err()
}

func (d *Dict) show(db *sqlx.DB) {
	var userName string
	fmt.Print("Enter username:\n")
	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Warnln(err)
		return
	}
	rows, err := db.Query("SELECT phone, description FROM accounts WHERE name=$1", userName)
	if err != nil {
		log.Warnln(err)
	}
	var userPhone string
	var userDesc string
	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&userPhone, &userDesc)
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
		fmt.Printf("%s's description: %s\n", userName, userDesc)
	}
	err = rows.Err()
}

func (d *Dict) find(db *sqlx.DB) {
	var userPhone string
	fmt.Print("Enter phone number:\n")
	_, err := fmt.Scan(&userPhone)
	if err != nil {
		log.Warnln(err)
		return
	}
	row := db.QueryRow("SELECT name FROM accounts WHERE phone=$1", userPhone)
	var userName string
	err = row.Scan(&userName)
	if err == nil {
		fmt.Printf("%s's phone: %s\n", userName, userPhone)
	} else {
		fmt.Println("Not found")
	}
}

func main() {
	db, err := NewPostgresDB()
	tx := db.MustBegin()
	if err != nil {
		log.Fatal("Cant create", err)
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
			dict.all(db)
		case command.Phone:
			dict.phone(db)
		case command.Desc:
			dict.desc(db)
		case command.Find:
			dict.find(db)
		case command.Show:
			dict.show(db)
		case command.Exit:
			return
		default:
			fmt.Printf("%s for help\n", command.Help)
		}
	}
}
