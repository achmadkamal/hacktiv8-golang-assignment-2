package controller

import "github.com/gin-gonic/gin"

type OrderController interface {
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	GetOrderById(c *gin.Context)
	UpdateOrderById(c *gin.Context)
	DeleteOrderById(c *gin.Context)
}
