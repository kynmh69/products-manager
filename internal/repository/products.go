package repository

import (
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
	result := p.db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *productRepository) GetAllProducts() ([]*domain.Products, error) {
	var products []*domain.Products
	if result := p.db.Find(products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
