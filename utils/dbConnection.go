package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormDbConnection() (*gorm.DB, *errs.AppError) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_ADDR"], myEnv["DB_PORT"], myEnv["DB_NAME"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errs.NewUnexpectError(err.Error())
	}
	return db, nil
}
func SqlxDbConnection() (*sqlx.DB, *errs.AppError) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_ADDR"], myEnv["DB_PORT"], myEnv["DB_NAME"])

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, errs.NewUnexpectError(err.Error())
	}
	return db, nil

}
