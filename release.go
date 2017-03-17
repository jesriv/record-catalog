package main

import "github.com/jinzhu/gorm"

type Release struct {
	gorm.Model
	Title		string	`json:"title"`
	CatalogId	string	`json:"catalog-id"`
}

type Releases []Release

func GetReleases() interface{} {
	db := Database()
	defer db.Close()

	var releases []Release
	results := db.Find(&releases)

	return results.Value
}

func (r *Release) Create() interface{} {
	db := Database()
	defer db.Close()

	return db.Create(r)
}