package database

import (
	"log"

	"github.com/kynmh69/products-manager/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	config := configs.NewMySQLConfig()
	var (
		db  *gorm.DB
		err error
	)

	if db, err = gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	return db
}
