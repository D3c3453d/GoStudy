package entity

type Account struct {
	Id       int    `db:"id" json:"-"`
	Name     string `db:"name" json:"name" binding:"required"`
	Password string `db:"password_hash" json:"password" binding:"required"`
	Phone    string `db:"phone" json:"phone" binding:"required"`
	Desc     string `db:"description" json:"description"`
}
