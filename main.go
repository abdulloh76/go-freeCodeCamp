package main

import (
	"fibercrm/database"
	"fibercrm/lead"
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DB_NAME  = "postgres"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/lead", lead.GetLeads)
	app.Get("api/lead/:id", lead.GetLead)
	app.Post("api/lead", lead.NewLead)
	app.Delete("api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", HOST, USER, PASSWORD, DB_NAME, PORT)
	var err error
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	fmt.Println("connection is opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database automigrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	// defer database.DBConn.Close()
}
