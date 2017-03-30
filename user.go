package main

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

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

func (u *User) Authenticate() (interface{}, error) {
	db := Database()
	defer db.Close()

	var user User

	result := db.Where("username = ?", u.Username).First(&user).Value
	hashedPassword := result.(*User).Password
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password));
	
	if err != nil {
		return nil, err
	}

	return result, err
}

func (u *User) Create() interface{} {
	db := Database()
	defer db.Close()

	return db.Create(u)
}

func (u *User) BeforeCreate() (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Password = string(hashedPassword)
	
	return
}
