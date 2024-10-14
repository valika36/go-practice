package Models

type Customer struct {
	Id    string `gorm:"type:char(36);primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *Customer) TableName() string {
	return "customer"
}
