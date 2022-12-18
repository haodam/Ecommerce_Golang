package db

import (
	accountmodel "Ecommerce_Golang/module/account/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {

	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)

	db.AutoMigrate(&accountmodel.Account{})

	return db
}
