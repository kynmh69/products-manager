package env

import (
	"log"
	"os"
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
