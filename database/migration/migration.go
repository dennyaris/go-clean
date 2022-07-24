package migration

import (
	"fmt"
	"go-clean/database"
	"go-clean/models"
	"log"
)

func RunMigration() {
	if err := database.DB.AutoMigrate(&models.Product{}); err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrate")
}
