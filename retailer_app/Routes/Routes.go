package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valikak/go-practice/retailer_app/Controllers"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/retailer-app")
	{
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product", Controllers.GetProducts)
		grp1.GET("product/:id", Controllers.GetProductById)
		grp1.PATCH("product/:id", Controllers.PatchProduct)
		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.GetOrderById)
		grp1.GET("order/history/:customerId", Controllers.GetCustomerOrders)
		grp1.GET("transactions", Controllers.GetAllTransactions)
	}
	return r
}
