package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// configure postgres for the db
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

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
