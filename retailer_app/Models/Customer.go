package Models

import "github.com/valikak/go-practice/retailer_app/Config"

func CreateCustomer(customer *Customer) (err error) {
	if err := Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
