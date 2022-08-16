package empdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	Email        string `sql:"unique"`
	ManagerID    *uint
	Manager      *Employee `gorm:"foreignkey:ManagerID;association_foreignkey:ID"`
	DepartmentID int
	Department   Department
}

type Department struct {
	gorm.Model
	Name      string
	Employees []Employee
}

func (e *Employee) AfterCreate() {
	fmt.Println("Sending email to ", e.Email)
}
