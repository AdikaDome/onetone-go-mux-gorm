package main

import (
	"onetoone/controller"
	"onetoone/databaseconnection"
)

func main() {
	databaseconnection.DatabaseConnection()
	controller.Controller()
}
