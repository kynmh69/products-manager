package domain

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name        string  `gorm:"name;primaryKey;varchar(255);not null" json:"name"`
	UnitPrice   float32 `gorm:"price;float(4);not null" json:"price"`
	Description string  `gorm:"description;varchar(255)" json:"description"`
	Category    string  `gorm:"size:50" json:"category"`
}
