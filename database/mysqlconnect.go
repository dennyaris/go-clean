package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Dbconnect() *gorm.DB {
	// dbHost := config.GetEnv("DB_HOST")
	// dbPort := config.GetEnv("DB_PORT")
	// dbUser := config.GetEnv("DB_USER")
	// dbPassword := config.GetEnv("DB_PASSWORD")
	// dbDatabase := config.GetEnv("DB_DATABASE")
	// socket := config.GetEnv("DB_SOCKET")

	// db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@unix("+socket+")/"+dbDatabase+"?parseTime=True") //localhost

	//postgree
	var err error

	database_url := os.Getenv("DATABASE_URL")

	if database_url == "" {
		database_url = "postgresql://postgres:@localhost:5432/go_clean?ssl_mode=disable&TimeZone=Asia/Jakarta"
	}

	dsn := database_url
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	return DB
}
