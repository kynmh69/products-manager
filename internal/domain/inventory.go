package domain

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID   uint     `gorm:"not null"`
	Product     Products `gorm:"foreignKey:ProductID"`
	Quantity    int      `gorm:"not null"`
	Location    string   `gorm:"size:50"`
	LastUpdated time.Time
}
