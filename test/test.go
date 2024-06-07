package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ginChart/model"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchart?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// 迁移 schema
	err = db.AutoMigrate(&model.UserBasic{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// 创建用户
	user := model.UserBasic{Name: "Alice", PassWord: "password123", Phone: "1234567890", Email: "alice@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println("User created successfully")
}