package app

import (
	"github.com/joho/godotenv"
	"github.com/pauliusluksys/golang-Reddit/domain"
	"github.com/pauliusluksys/golang-Reddit/migrations"
	"github.com/pauliusluksys/golang-Reddit/seeds"
	"log"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	//db := domain.GormDbConnections()
	dbSqlx := domain.SqlxDbConnections()
	migrations.PostMigration()
	seeds.Execute(dbSqlx, "PostSeed")
	domain.PostGorm{}.TableName()

	//r := routes()
	//
	//err = r.Run(":8080")
	//if err != nil {
	//	panic("Gin routing has failed: " + err.Error())
	//}
}
