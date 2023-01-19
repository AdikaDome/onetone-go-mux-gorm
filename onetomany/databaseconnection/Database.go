package databaseconnection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"onetomany/model"
)

var dsn = "host=localhost user=postgres password=root dbname=onetomanygo port=5432 sslmode=disable"

func Database() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		log.Println("Database connected")
	}
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(model.Employee{}, model.Department{})
	return db

}
