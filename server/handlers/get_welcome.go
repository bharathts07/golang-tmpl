package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/internal/models"
	"github.com/bharathts07/pokke/server/middleware"
)

func GetWelcomeComponents(c *gin.Context) {
	db := middleware.GetDB(c)
	data, _ := db.GetComponents(c, "welcome", []models.ComponentType{})
	c.JSON(200, data)
}
