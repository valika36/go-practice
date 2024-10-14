package Models

import (
	"time"
)

type StatusType string

const (
	OrderPlaced StatusType = "PLACED"
	Processed   StatusType = "PROCESSED"
	Failed      StatusType = "FAILED"
)

type Order struct {
	Id          string     `gorm:"type:char(36);primaryKey"`
	CustomerId  string     `json:"customerId"`
	Customer    Customer   `gorm:"foreignKey:CustomerId;references:Id"`
	ProductId   string     `json:"productId"`
	Product     Product    `gorm:"foreignKey:ProductId;references:Id"`
	Quantity    uint       `json:"quantity"`
	TotalPrice  uint       `json:"totalPrice"`
	Status      StatusType `json:"status"`
	CreatedTime time.Time  `json:"createdTime"`
}

func (o *Order) TableName() string {
	return "order"
}
