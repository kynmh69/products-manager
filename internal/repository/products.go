package repository

import (
	"github.com/kynmh69/products-manager/consts"
	"github.com/kynmh69/products-manager/env"
	"github.com/kynmh69/products-manager/internal/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (p *productRepository) CreateProduct(product *domain.Products) error {
	var result *gorm.DB
	if env.GetGinMode() == consts.PRODUCTION {
		result = p.db.Create(product)
	} else {
		result = p.db.Debug().Create(product)
	}
	
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *productRepository) GetAllProducts() ([]*domain.Products, error) {
	var products []*domain.Products
	db := p.db.Debug()
	if env.GetGinMode() == consts.PRODUCTION {
		db = p.db
	}
	if result := db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
