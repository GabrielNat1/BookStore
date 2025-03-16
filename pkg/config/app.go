package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
