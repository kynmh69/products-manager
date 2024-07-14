package main

import (
	"github.com/kynmh69/products-manager/internal/app/router"
)

func main() {
	r := router.NewRouter()
	r.Engine.Run()
}
