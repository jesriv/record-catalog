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

func GetUsers() interface{} {
	db := Database()
	defer db.Close()

	var users []User

	return db.Find(&users).Value
}

func (u *User) Authenticate() interface{} {
	db := Database()
	defer db.Close()

	var users []User

	return db.Where("username = ? AND password = ?", u.Username, u.Password).Find(&users).Value
}

func (u *User) Create() interface{} {
	db := Database()
	defer db.Close()

	return db.Create(u)
}

