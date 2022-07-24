package controller

import (
	"go-clean/database"
	"go-clean/models"
	"go-clean/present"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var (
		data []models.Product
	)

	db := database.Dbconnect()
	db.Find(&data)

	return c.JSON(present.Success("Success", data, nil))
}

func GetById(c *fiber.Ctx) error {
	var (
		data models.Product
	)

	prodId := c.Params("id")
	db := database.Dbconnect()
	if err := db.Find(&data, prodId); err != nil {
		c.Status(400)
		return c.JSON(present.Error("Data not found", nil))
	}

	return c.JSON(present.Success("Success", data, nil))
}

func Save(c *fiber.Ctx) error {
	db := database.Dbconnect()
	data := new(models.Product)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	if validateIfExist(data.Name) == true {
		db.Create(&data)
		return c.JSON(present.Success("Success", data, nil))
	} else {
		c.Status(401)
		return c.JSON(present.Error("Data "+data.Name+" already exist", nil))
	}
}

func UpdateById(c *fiber.Ctx) error {
	dataReq := new(models.Product)
	if err := c.BodyParser(dataReq); err != nil {
		c.Status(400)
		return c.JSON(present.Error("Bad request update", nil))
	}

	var data models.Product
	prodId := c.Params("id")

	db := database.Dbconnect()
	err := db.Find(&data, prodId).Error

	if err != nil {
		c.Status(400)
		return c.JSON(present.Error("Data not found", nil))
	} else {
		data.Name = dataReq.Name
		data.Description = dataReq.Description

		errUpdate := db.Save(&data).Error
		if errUpdate != nil {
			c.Status(500)
			return c.JSON(present.Error("Update to database error", nil))
		}
	}

	return c.JSON(present.Success("Update data success", dataReq, nil))
}

func Delete(c *fiber.Ctx) error {
	var (
		data models.Product
	)

	prodId := c.Params("id")
	db := database.Dbconnect()
	db.Delete(&data, prodId)

	return c.JSON(present.Success("Delete Product Success", nil, nil))
}

func validateIfExist(productName string) bool {
	var tmp models.Product
	db := database.Dbconnect()
	err := db.Table("product").Where("name = ?", productName).First(&tmp)

	if err.RowsAffected == 0 {
		return true
	} else {
		return false
	}
}
