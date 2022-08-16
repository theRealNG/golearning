package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/theRealNG/golearning/emp/empdb"
)

func main() {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Sprintf("Failed to open connection to db due to: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&empdb.Employee{}, &empdb.Department{})
	empdb.SeedDb(db)
	var employees []empdb.Employee
	db.Debug().Preload("Department").Preload("Manager").Find(&employees)
	for _, emp := range employees {
		fmt.Printf("Employee Details: %v, Manager Details: %v, Department: %v \n", emp, emp.Manager, emp.Department)
	}
}
