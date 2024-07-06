package configs

import (
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kynmh69/products-manager/consts"
	"github.com/kynmh69/products-manager/env"
)

func NewMySQLConfig() *mysql.Config {
	username := env.Get(consts.DB_USER)
	password := env.Get(consts.DB_PASS)
	addr := fmt.Sprintf("%s:%s", env.Get(consts.DB_HOST), env.Get(consts.DB_PORT))
	dbName := env.Get(consts.DB_NAME)
	loc := getLocation()
	return &mysql.Config{
		User:      username,
		Passwd:    password,
		Net:       consts.PROTOCOL,
		Addr:      addr,
		DBName:    dbName,
		Loc:       loc,
		ParseTime: true,
	}
}

func getLocation() *time.Location {
	var loc *time.Location
	if l, err := time.LoadLocation(env.Get("TZ")); err != nil {
		log.Fatalln(err)
	} else {
		loc = l
	}
	return loc
}
