package router

import (
	"hacktiv8_golang_assignment_2/controller"
	"hacktiv8_golang_assignment_2/database"
	"hacktiv8_golang_assignment_2/repository"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	orderRepo := repository.NewOrderRepository(database.DB)

	orderCtrl := controller.NewOrderController(orderRepo)

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/", orderCtrl.CreateOrder)
		orderRoutes.GET("/", orderCtrl.GetOrders)
		orderRoutes.GET("/:id", orderCtrl.GetOrderById)
		orderRoutes.PUT("/:id", orderCtrl.UpdateOrderById)
		orderRoutes.DELETE("/:id", orderCtrl.DeleteOrderById)
	}

	return r
}
