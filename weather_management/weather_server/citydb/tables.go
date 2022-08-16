package citydb

import (
	"github.com/jinzhu/gorm"
)

type City struct {
	gorm.Model
	Lat  float32
	Long float32
	Name string `gorm:"unique"`
}
