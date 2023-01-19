package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName    string       `json:"EmpName"`
	EmpSalary  float32      `json:"EmpSalary"`
	Department []Department `json:"Department"`
}
