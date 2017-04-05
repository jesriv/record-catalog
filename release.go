package main

import (
	"github.com/jinzhu/gorm"
)

type Release struct {
	gorm.Model
	Title		string	`json:"title"`
	CatalogId	string	`json:"catalog_id"`
}

type Releases []Release

func GetReleases() interface{} {
	db := Database()
	defer db.Close()

	var releases []Release

	return db.Find(&releases).Value
}

func (r *Release) Get(id string) interface{} {
	db := Database()
	defer db.Close()

	return db.Find(r, id).Value
}

func (r *Release) Create() interface{} {
	db := Database()
	defer db.Close()

	return db.Create(r)
}

func (r *Release) Update(id string) interface{} {
	db := Database()
	defer db.Close()

	return db.Model(r).Where("id = ?", id).Updates(map[string]interface{}{
			"title"			: r.Title,
			"catalog_id"	: r.CatalogId,
		})
}