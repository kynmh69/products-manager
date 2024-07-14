package database

import (
	"log"

	"github.com/kynmh69/products-manager/configs"
	"github.com/kynmh69/products-manager/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	logging.NewLogger()
	NewGorm()
}

func NewGorm() *gorm.DB {
	suger := logging.NewLogger()
	if db != nil {
		return db
	}
	config := configs.NewMySQLConfig()
	var err error
	if db, err = gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	suger.Infoln("Connected db.")
	return db
}
