package main

import (
	"log"

	"github.com/kynmh69/products-manager/database"
	"github.com/kynmh69/products-manager/internal/domain"
)

func main() {
	db := database.NewGorm()
	if err := db.AutoMigrate(&domain.Products{}, &domain.Inventory{}); err != nil {
		log.Fatalln(err)
	}
}
