package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/products-manager/internal/app/handler"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := &Router{Engine: gin.Default()}
	router.initHandler()
	return router
}

func (r *Router) initHandler() {
	r.Engine.POST("/products", handler.CreateProducts())
	r.Engine.GET("/products", handler.GetAllProducts())
}
