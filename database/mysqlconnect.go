package database

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
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

	// db, err := gorm.Open("postgres", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+"?parseTime=True")
	// db, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@unix("+socket+")/"+dbDatabase+"?parseTime=True") //localhost

	var err error
	//postgree
	const POSTGREESQL = "postgres://fptywazsibmxxb:4328999b1a1643b930c87d201fd7db3926a30f2b1c29f59eb6816ae5ad37c93e@ec2-44-199-143-43.compute-1.amazonaws.com:5432/d82o9n2p4hum9k"
	dsn := POSTGREESQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	return DB
}
