package database

import (
	"fmt"

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
	const POSTGREESQL = "postgres://wzclpugyoxjxuk:59121310f6e4cfc0e832a864a3319873f0fabe074b5362bb7cc89ba15bf3d4a6@ec2-54-208-104-27.compute-1.amazonaws.com:5432/d919ovmafh5m6f"
	dsn := POSTGREESQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	return DB
}
