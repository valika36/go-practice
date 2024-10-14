package Models

import "github.com/valikak/go-practice/retailer_app/Config"

func CreateProduct(product *Product) (err error) {
	if err := Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProducts(product *[]Product) (err error) {
	if err := Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Product, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product) (err error) {
	if err := Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
