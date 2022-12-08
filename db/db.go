package db

import (
	"Ecommerce_Golang/module/account/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {

	dsn := "host=localhost user=root password=ecommerce123 dbname=ecommercegolang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)

	db.AutoMigrate(&model.Account{})

	return db
}
