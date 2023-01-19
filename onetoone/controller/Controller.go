package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"onetoone/service"
)

func Controller() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", service.PostEmployee).Methods("POST")
	r.HandleFunc("/employees/{employeeId}", service.GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", service.GetAllEmployee).Methods("Get")
	r.HandleFunc("/employees/{employeeId}", service.UpdateEmployee).Methods("Update")
	r.HandleFunc("/employees/{employeeId}", service.DeleteEmployeeById).Methods("Delete")
	http.ListenAndServe(":8081", r)
}
