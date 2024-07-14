package main

import (
	"time"

	"github.com/kynmh69/products-manager/consts"
	"github.com/kynmh69/products-manager/database"
	"github.com/kynmh69/products-manager/env"
	"github.com/kynmh69/products-manager/internal/domain"
	"github.com/kynmh69/products-manager/logging"
)

func main() {
	db := database.NewGorm()
	logger := logging.NewLogger()
	if env.GetGinMode() != consts.PRODUCTION {
		db = db.Debug()
	}
	products := domain.Products{
		Name:      "orange",
		UnitPrice: 180,
		Category:  consts.FRUIT,
	}
	inventory := domain.Inventory{
		ProductID:   1,
		Quantity:    100,
		LastUpdated: time.Now(),
	}
	if result := db.FirstOrCreate(&products); result.Error != nil {
		logger.Fatalln(result.Error)
	}
	if result := db.FirstOrCreate(&inventory); result.Error != nil {
		logger.Fatalln(result.Error)
	}
}
