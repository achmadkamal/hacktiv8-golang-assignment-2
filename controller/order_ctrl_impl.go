package controller

import (
	"hacktiv8_golang_assignment_2/entity"
	"hacktiv8_golang_assignment_2/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type orderControllerImpl struct {
	orderRepo repository.OrderRepository
}

func NewOrderController(orderRepo repository.OrderRepository) OrderController {
	return &orderControllerImpl{orderRepo: orderRepo}
}

func (o *orderControllerImpl) CreateOrder(c *gin.Context) {
	var orderRequest struct {
		CustomerName string `json:"customer_name" binding:"required"`
		OrderedAt    string `json:"ordered_at" binding:"required"`
		Items        []struct {
			ItemCode    string `json:"item_code" binding:"required"`
			Description string `json:"description" binding:"required"`
			Quantity    uint   `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderedAt, err := time.Parse(time.RFC3339, orderRequest.OrderedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	order := entity.Order{
		CustomerName: orderRequest.CustomerName,
		OrderedAt:    orderedAt,
		Items:        []entity.Item{},
	}

	for _, item := range orderRequest.Items {
		order.Items = append(order.Items, entity.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	if err := o.orderRepo.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (o *orderControllerImpl) GetOrders(c *gin.Context) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (o *orderControllerImpl) GetOrderById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := o.orderRepo.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (o *orderControllerImpl) UpdateOrderById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var orderRequest struct {
		CustomerName string `json:"customer_name" binding:"required"`
		Items        []struct {
			ItemCode    string `json:"item_code" binding:"required"`
			Description string `json:"description" binding:"required"`
			Quantity    uint   `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := o.orderRepo.GetOrderById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.CustomerName = orderRequest.CustomerName
	order.Items = []entity.Item{}

	for _, item := range orderRequest.Items {
		order.Items = append(order.Items, entity.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	if err := o.orderRepo.UpdateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (o *orderControllerImpl) DeleteOrderById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := o.orderRepo.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
