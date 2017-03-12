package main

type Release struct {
	Title		string	`json:"title"`
	CatalogId	string	`json:"catalog-id"`
}

type Releases []Release