package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Product struct {
	gorm.Model
	Name string
	Price uint
}

func OpenDatabase(database_path string)(*gorm.DB) {
	/* Open the Database */
	db, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	
	return db
}

func GetDatabase(db *gorm.DB)([]Product) {
	/* Get all the entries of the Database */
	var products []Product // products list
	db.Find(&products) // find entries of products database

	return products
}
