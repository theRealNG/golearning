package citydb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type CityDB struct {
	connection *gorm.DB
}

func (c *CityDB) Conn() {
	db, err := gorm.Open("sqlite3", "city.db")
	if err != nil {
		fmt.Printf("Failed to connect to DB: %v", err)
	}
	c.connection = db
}

func (c *CityDB) CloseConn() {
	c.connection.Close()
}

func (c *CityDB) SetupDB() {
	if c.connection == nil {
		c.Conn()
	}
	c.connection.AutoMigrate(&City{})
}

func (c *CityDB) CreateCity(city *City) {
	if c.connection == nil {
		c.Conn()
	}
	c.connection.Debug().Create(&city)
}

func (c *CityDB) FindCity(city *City) {
	if c.connection == nil {
		c.Conn()
	}
	c.connection.Debug().Where("name = ?", city.Name).First(&city)
}
