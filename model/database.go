package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("share-fridge.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&User{}, &Food{}, &Fridge{})
}
