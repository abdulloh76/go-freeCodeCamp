package config

import (
	"book-store/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// configure postgres for the db
	host := utils.POSTGRES_HOST
	user := utils.POSTGRES_USER
	password := utils.POSTGRES_PASSWORD
	dbname := utils.POSTGRES_DB
	port := utils.POSTGRES_PORT

	dsn := "host=" + host + "user=" + user + "password=" + password + "dbname=" + dbname + "port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
