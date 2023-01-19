package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"onetoone/databaseconnection"
	"onetoone/model"
)

func PostEmployee(response http.ResponseWriter, request *http.Request) {
	log.Printf("adding new employee.....")
	response.Header().Set("content-type", "application/json")
	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)
	db := databaseconnection.DatabaseConnection()
	db.Create(&employee)
	json.NewEncoder(response).Encode(employee)
	return

}
func GetEmployeeById(response http.ResponseWriter, request *http.Request) {
	log.Printf("get employee information")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.DatabaseConnection()
	params := mux.Vars(request)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Department").First(&employee, params["employeeId"])
	json.NewEncoder(response).Encode(employee)
	return

}

func GetAllEmployee(response http.ResponseWriter, request *http.Request) {
	log.Printf("get all employees")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.DatabaseConnection()
	var employees []model.Employee
	db.Find(&employees)
	json.NewEncoder(response).Encode(employees)
	return

}
func UpdateEmployee(response http.ResponseWriter, request *http.Request) {
	log.Printf("Updating employee information")
	response.Header().Set("content-type", "apllication/json")
	db := databaseconnection.DatabaseConnection()
	var params = mux.Vars(request)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Department").First(&employee, params["employeeId"])
	json.NewEncoder(response).Encode(employee)
	db.Save(&employee)
	return
}
func DeleteEmployeeById(response http.ResponseWriter, request *http.Request) {
	log.Print("deleting employee information.....")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.DatabaseConnection()
	params := mux.Vars(request)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Department").First(&employee, params["employeeId"])
	db.Delete(&employee)
	json.NewEncoder(response).Encode("delete successfully")
	return

}
