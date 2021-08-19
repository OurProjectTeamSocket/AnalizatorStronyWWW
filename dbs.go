package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Get(db * gorm.DB, name string, object Website) *Website {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	result := db.Table(name).Last(&object)

	var item Website

	result.First(&item)

	return &item
}