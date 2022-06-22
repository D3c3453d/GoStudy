package main

import "fmt"

func help(dict map[string]map[string]string) {
	fmt.Println("/add to add new account")
	fmt.Println("/all to see all accounts")
	fmt.Println("/desc to see description of the account")
	fmt.Println("/phone to see phone of the account")
	fmt.Println("/find to find by phone number")
}

func add(dict map[string]map[string]string) {
	var userName string
	var userPhone string
	var userDesc string

	fmt.Print("Enter your username: ")
	fmt.Scanln(&userName)
	fmt.Print("Enter your phone number: ")
	fmt.Scanln(&userPhone)
	fmt.Print("Enter your description: ")
	fmt.Scanln(&userDesc)
	dict[userName] = map[string]string{}
	dict[userName]["Phone"] = userPhone
	dict[userName]["Description"] = userDesc
}

func all(dict map[string]map[string]string) {
	for userName := range dict {
		fmt.Println(userName)
	}
}

func phone(dict map[string]map[string]string) {
	var userName string
	fmt.Print("Enter username: ")
	fmt.Scanln(&userName)
	fmt.Printf("%s's phone number: %s\n", userName, dict[userName]["Phone"])
}

func desc(dict map[string]map[string]string) {
	var userName string
	fmt.Print("Enter username: ")
	fmt.Scanln(&userName)
	fmt.Printf("%s's description: %s\n", userName, dict[userName]["Description"])
}

func find(dict map[string]map[string]string) {
	var userPhone string
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&userPhone)
	for userName := range dict {
		if userPhone == dict[userName]["Phone"] {
			fmt.Printf("%s's phone number: %s\n", userName, dict[userName]["Phone"])
			return
		}
	}
	fmt.Println("Not found")
}

func main() {
	dict := map[string]map[string]string{}
	var input string
	for {
		fmt.Scanln(&input)
		switch input {
		case "/help":
			help(dict)
		case "/add":
			add(dict)
		case "/all":
			all(dict)
		case "/phone":
			phone(dict)
		case "/desc":
			desc(dict)
		case "/find":
			find(dict)
		case "/exit":
			break
		default:
			fmt.Println("/help for help")
		}
	}
}
