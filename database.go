package main

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	db, err := gorm.Open("mysql", username+":"+password+"@/release_catalog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Release{})

	return db
}