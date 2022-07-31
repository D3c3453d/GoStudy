package main

type Account struct {
	userId    string `db:"id"`
	userName  string `db:"name"`
	userPhone string `db:"phone"`
	userDesc  string `db:"description"`
}
