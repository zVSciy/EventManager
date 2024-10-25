package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPayment := services.CreatePayment(payment)
	c.JSON(http.StatusCreated, createdPayment)
}

// func ProcessPayment(c *gin.Context) {

// }
