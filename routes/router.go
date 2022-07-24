package routes

import (
	"go-clean/controller"

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

	_ = r.Listen(":90")
}
