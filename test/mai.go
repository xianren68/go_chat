package main

import (
	"go_chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:010825lwj@tcp(47.92.99.157:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移数据库
	err = db.AutoMigrate(&models.UserBasic{}, &models.Relation{})
	if err != nil {
		panic(err)
	}

}
