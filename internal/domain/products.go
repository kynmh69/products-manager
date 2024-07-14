package domain

type Products struct {
	Model
	Name        string  `gorm:"name;unique;size:255;not null" json:"name" binding:"required"`
	UnitPrice   float32 `gorm:"price;float(4);not null" json:"price" binding:"required"`
	Description string  `gorm:"description;size:255" json:"description"`
	Category    string  `gorm:"size:50" json:"category"`
}
