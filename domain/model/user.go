package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"user_name" gorm:"varchar(16);not_null"`
	Password string `json:"password" gorm:"varchar(16);not_null"`
}
