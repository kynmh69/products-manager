package env

import (
	"log"
	"os"

	"github.com/kynmh69/products-manager/consts"
)

func Get(key string) string {
	var value string
	if v, ok := os.LookupEnv(key); ok {
		value = v
	} else {
		log.Fatalln("don't look up env", key)
	}
	return value
}

func GetGinMode() string {
	return os.Getenv(consts.GIN_MODE)
}
