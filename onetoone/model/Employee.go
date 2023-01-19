package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeName string     `json:"employeeName"`
	Salary       float32    `json:"salary"`
	Department   Department `json:"department"`
	DepartmentId uint       `gorm:"foreignKey:ID" json:"_"`
}
