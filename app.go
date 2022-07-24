package main

import (
	"go-clean/database"
	"go-clean/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = os.Setenv("TZ", "Asia/Jakarta")

	database.Dbconnect()

	app := fiber.New()
	routes.RouterInit(app)
}
