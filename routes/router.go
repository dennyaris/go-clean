package routes

import (
	"go-clean/controller"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RouterInit(r *fiber.App) {
	r.Use(cors.New())

	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "GO Fiber + GORM",
		})
	})

	p := r.Group("/product")
	p.Get("/", controller.GetAll)
	p.Get("/:id", controller.GetById)
	p.Post("/", controller.Save)
	p.Put("/:id", controller.UpdateById)
	p.Delete("/:id", controller.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	appList := r.Listen(":" + port)
	if appList != nil {
		log.Println("fail to listen go fiber server")
		os.Exit(1)
	}
}
