package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string `json:"DepartmentName"`
	EmployeeId     uint   `gorm:"foreignKey:ID" json:"_"`
}
