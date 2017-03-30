package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type Users []User

func (u *User) Authenticate() interface{} {
	db := Database()
	defer db.Close()

	var users []User

	return db.Where("username = ? AND password = ?", u.Username, u.Password).Find(&users).Value
}