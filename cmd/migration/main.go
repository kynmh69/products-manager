package main

import (
	"log"

	"github.com/kynmh69/products-manager/consts"
	"github.com/kynmh69/products-manager/database"
	"github.com/kynmh69/products-manager/env"
	"github.com/kynmh69/products-manager/internal/domain"
)

func main() {
	db := database.NewGorm()

	if env.GetGinMode() != consts.PRODUCTION {
		db = db.Debug()
	}
	if err := db.AutoMigrate(&domain.Products{}, &domain.Inventory{}); err != nil {
		log.Fatalln(err)
	}
}
