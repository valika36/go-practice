package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valikak/go-practice/retailer_app/Models"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var order Models.OrderRequest
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderById(&order, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetCustomerOrders(c *gin.Context) {
	id := c.Params.ByName("customerId")
	var order []Models.Order
	err := Models.GetOrderByCustomerId(&order, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetAllTransactions(c *gin.Context) {
	var transactions []Models.TransactionResponse
	err := Models.GetAllTransactions(&transactions)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transactions)
	}
}
