package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conect() {
	dsn := "db_app:db_app@tcp(localhost:8002)/db_app?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := sql.Open("mysql", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error connect database")
		log.Panic(DB)

	}
}
