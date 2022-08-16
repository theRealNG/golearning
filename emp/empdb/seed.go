package empdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SeedDb(db *gorm.DB) {
	departments := make([]Department, 0, 10)
	departments = append(departments, Department{Name: "IT Dept"})
	for _, dep := range departments {
		db.Debug().Create(&dep)
	}

	department := Department{}
	db.First(&department)
	// employees := make([]Employee, 0, 10)
	var employees []Employee
	employees = append(employees, Employee{Name: "Manager", Email: "test@manager.com", Department: department})
	employees = append(employees, Employee{Name: "Test", Email: "test@test.com", Department: department})
	employees = append(employees, Employee{Name: "Test 1", Email: "test1@test.com", Department: department})
	for _, emp := range employees {
		result := db.Debug().Create(&emp)
		if result.Error != nil {
			fmt.Println("Failed to create employee due to: %v", result.Error)
		}
	}

	// Retrieve Manager
	var manager Employee
	db.Debug().Where(&Employee{Email: "test@manager.com"}).Find(&manager)
	// Update manager ID
	db.Debug().Not(&Employee{Email: "test@manager.com"}).Find(&employees).Update("manager_id", manager.ID)
}
