package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func ConnectDB() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(&User{})
	fmt.Println("Connect db")
	return database
}

func main() {
	var db *gorm.DB = ConnectDB()
	user1 := User{Username: "suntipap", Password: "1212"}
	db.Create(&user1)
}
