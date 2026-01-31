package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func h_user_selection(c *gin.Context) {

	var jasonData struct {
		Monto        string `json:"monto" binding:"required"`
		Selection    string `json:"selection" binding:"required"`
		Plazo        string `json:"plazo" binding:"required"`
		Periodicidad string `json:"periodicidad" binding:"required"`
	}
	err := c.ShouldBindJSON(&jasonData)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	// logic for saving user selection
	c.JSON(200, gin.H{
		"message": "Selection saved",
	})
}
func h_user_info(c *gin.Context) {
	token := c.Query("token")
	// logic for fetching user info by token
	if token == "valid-token" {
		c.JSON(200, gin.H{
			"user":            "Francisco Marquez",
			"empleador":       "Restaurante La Sandia",
			"ingreso_mensual": 15000,
			"monto_aprobado":  5000,
			"monto_minimo":    1000,
			"monto_maximo":    5000,
		})
	} else {
		c.JSON(401, gin.H{
			"error": "Invalid token",
		})
	}
}

func main() {
	router := gin.Default()
	router.GET("/user-info", h_user_info)
	router.POST("/user-selection", h_user_selection)
	router.Run() // listens on 0.0.0.0:8080 by default
}
