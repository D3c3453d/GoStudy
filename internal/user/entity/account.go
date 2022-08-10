package entity

type Account struct {
	UserId    string `db:"id" json:"-"`
	UserName  string `db:"name" json:"name"`
	UserPhone string `db:"phone" json:"phone"`
	UserDesc  string `db:"description" json:"description"`
}
