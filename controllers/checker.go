package adds

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func check() {
	db, err := gorm.Open(sqlite.Open("db.db"), gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	result := db.Table("users").Find(&User{})

	for x, y := range result.Row() {



	}
}
