package domain

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GormDbConnections() *gorm.DB {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_ADDR"], myEnv["DB_PORT"], myEnv["DB_NAME"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func SqlxDbConnections() *sqlx.DB {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_ADDR"], myEnv["DB_PORT"], myEnv["DB_NAME"])

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db

}
