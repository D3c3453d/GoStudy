package model

type Account struct {
	UserId    string `db:"id"`
	UserName  string `db:"name"`
	UserPhone string `db:"phone"`
	UserDesc  string `db:"description"`
}
