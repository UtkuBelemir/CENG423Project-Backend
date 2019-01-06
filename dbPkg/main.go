package dbPkg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const dbHost = "localhost"
const dbDialect = "host=" + dbHost + " user=aybu dbname=aybu_student_system sslmode=disable password=aybu123"

var err error

type dbConnector struct {
	DB *gorm.DB
}

var dbInstance *dbConnector

func New() *dbConnector {
	if dbInstance != nil && dbInstance.DB != nil {
		fmt.Println("If worked")
		return dbInstance
	} else {
		fmt.Println("Else worked")
		dbInstance = new(dbConnector)

		dbInstance.DB, err = gorm.Open("postgres", dbDialect)
		if err != nil {
			panic("Failed to create database connection on host : " + dbHost + ". Error : " + err.Error())
		}
		return dbInstance
	}
}
