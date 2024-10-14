package Models

type Product struct {
	Id          string `gorm:"type:char(36);primaryKey"`
	ProductName string `json:"productName"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

func (p *Product) TableName() string {
	return "product"
}
