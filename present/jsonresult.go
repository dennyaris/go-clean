package present

import "github.com/gofiber/fiber"

func Success(message string, data interface{}, meta interface{}) *fiber.Map {
	if meta != nil {
		return &fiber.Map{
			"data":    data,
			"meta":    meta,
			"message": message,
			"status":  1,
		}
	} else {
		return &fiber.Map{
			"data":    data,
			"message": message,
			"status":  1,
		}
	}
}

func Error(message string, data interface{}) *fiber.Map {
	return &fiber.Map{
		"data":    data,
		"message": message,
		"status":  0,
	}
}
