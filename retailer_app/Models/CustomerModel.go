package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	Id    string `gorm:"type:char(36);primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return
}

func (c *Customer) TableName() string {
	return "customer"
}
