package domain

import (
	"time"
)

type Inventory struct {
	Model
	ProductID   uint     `gorm:"not null" binding:"required"`
	Product     Products `gorm:"foreignKey:ProductID"`
	Quantity    int      `gorm:"not null" binding:"required"`
	Location    string   `gorm:"size:50"`
	LastUpdated time.Time
}
