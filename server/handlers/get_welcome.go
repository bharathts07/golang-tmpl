package handlers

import (
	"github.com/bharathts07/pokke/internal/models"
	"github.com/gin-gonic/gin"
)

func GetWelcome(c *gin.Context) {
	data := models.Welcome{
		Message:"ようこそう",
		Home:"https://0.0.0.0:8080/home",
	}
	c.JSON(200, data)
}
