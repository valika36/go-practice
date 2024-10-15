package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Id          string `gorm:"type:char(36);primaryKey" json:"id"`
	ProductName string `json:"productName" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.New().String()
	return
}

func (p *Product) TableName() string {
	return "product"
}
