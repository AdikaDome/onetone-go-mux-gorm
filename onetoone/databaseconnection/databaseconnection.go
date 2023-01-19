package databaseconnection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"onetoone/model"
)

var dsn = "host=localhost user=postgres password=root dbname=emp_depart port=5432 sslmode=disable"

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		log.Println("database connected")
	}
	if err != nil {
		log.Println(err)

	}
	db.AutoMigrate(model.Employee{}, model.Department{})
	return db

}
