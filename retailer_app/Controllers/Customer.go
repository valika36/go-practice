package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valikak/go-practice/retailer_app/Models"
	"net/http"
)

func CreateCustomer(c *gin.Context) {
	var customer Models.Customer
	c.BindJSON(&customer)
	err := Models.CreateCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
