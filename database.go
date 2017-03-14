package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/release_catalog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	return db
}

// type DBCaller interface {
// 	GetAll() interface{}
// }