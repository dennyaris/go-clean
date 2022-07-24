package database

import (
	"go-clean/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Dbconnect() *gorm.DB {
	// dbHost := config.GetEnv("DB_HOST")
	// dbPort := config.GetEnv("DB_PORT")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbDatabase := config.GetEnv("DB_DATABASE")
	socket := config.GetEnv("DB_SOCKET")

	// db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+"?parseTime=True")
	db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@unix("+socket+")/"+dbDatabase+"?parseTime=True")
	if err != nil {
		panic(err)
	}

	return db
}
