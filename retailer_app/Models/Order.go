package Models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/valikak/go-practice/retailer_app/Config"
	"time"
)

type OrderRequest struct {
	CustomerId string `json:"customerId" binding:"required"`
	ProductId  string `json:"productId" binding:"required"`
	Quantity   uint   `json:"quantity" binding:"required"`
}

func CreateOrder(order *OrderRequest) (err error) {

	canPlaceOrder, err := CanPlaceOrder(order.CustomerId)
	if err != nil {
		fmt.Println("Error while placing order")
		return nil
	}
	if !canPlaceOrder {
		return fmt.Errorf("cannot place order in cool-down period")
	}

	var product Product
	if err := Config.DB.First(&product, "id = ?", order.ProductId).Error; err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.Quantity < order.Quantity {
		return fmt.Errorf("insufficient product quantity")
	}

	totalPrice := product.Price * order.Quantity

	newOrder := Order{
		CustomerId: order.CustomerId,
		ProductId:  order.ProductId,
		Quantity:   order.Quantity,
		TotalPrice: totalPrice,
		Status:     "PLACED",
	}

	Config.DB.Create(&newOrder)

	product.Quantity -= order.Quantity
	if err := Config.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderById(order *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByCustomerId(order *[]Order, customerId string) (err error) {
	if err := Config.DB.Where("customerId = ?", customerId).Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTransactions(transactions *[]TransactionResponse) (err error) {
	var orders []Order
	if err := Config.DB.Preload("Customer").Preload("Product").Find(&orders).Error; err != nil {
		return err
	}

	for _, order := range orders {
		*transactions = append(*transactions, TransactionResponse{
			OrderID:    order.Id,
			CustomerId: order.Customer.Id,
			ProductId:  order.Product.Id,
			Quantity:   order.Quantity,
			TotalPrice: order.Product.Price * order.Quantity,
			Status:     string(order.Status),
		})
	}
	return nil

}

func GetLastOrderByCustomerId(order *Order, customerId string) (err error) {
	if err := Config.DB.Where("customer_id = ?", customerId).
		Order("created_time DESC").First(order).Error; err != nil {
		return err
	}
	return nil
}

func CanPlaceOrder(customerId string) (bool, error) {
	var lastOrder Order
	err := GetLastOrderByCustomerId(&lastOrder, customerId)

	if err != nil && err == gorm.ErrRecordNotFound {
		return true, nil
	} else if err != nil {
		return false, err
	}

	now := time.Now()
	coolDownEnd := lastOrder.CreatedTime.Add(5 * time.Minute)
	if now.Before(coolDownEnd) {
		return false, nil
	}

	return true, nil
}
