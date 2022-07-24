package main

import (
	"go-clean/database"
	"go-clean/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = os.Setenv("TZ", "Asia/Jakarta")

	db := database.Dbconnect()
	database.RunMigrate(db)

	app := fiber.New()
	routes.RouterInit(app)
}
