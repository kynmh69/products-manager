package database

import (
	"log"

	"github.com/kynmh69/products-manager/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	NewGorm()
}

func NewGorm() *gorm.DB {
	if db != nil {
		return db
	}
	config := configs.NewMySQLConfig()
	var err error
	if db, err = gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected db.")
	return db
}
