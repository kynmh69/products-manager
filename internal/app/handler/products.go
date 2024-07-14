package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kynmh69/products-manager/database"
	"github.com/kynmh69/products-manager/internal/domain"
	"github.com/kynmh69/products-manager/internal/repository"
)

func CreateProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Products
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		repo := repository.NewProductRepository(database.NewGorm())
		if err := repo.CreateProduct(&product); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"msg": err.Error()})
		}
		ctx.JSON(http.StatusCreated, product)
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repo := repository.NewProductRepository(database.NewGorm())
		if products, err := repo.GetAllProducts(); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, products)
		}
	}
}
