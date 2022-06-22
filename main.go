package main

import "fmt"

func main() {
	fmt.Println("/help for help")
	dict := map[string]map[string]string{}
	var input string
	for {
		fmt.Scanln(&input)
		if input == "/help" {
			fmt.Println("/add to add new account")
			fmt.Println("/all to see all accounts")
			fmt.Println("/desc to see description of the account")
			fmt.Println("/phone to see phone of the account")
		}
		if input == "/add" {
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
		if input == "/all" {
			for userName := range dict {
				fmt.Println(userName)
			}
		}
	}
}
