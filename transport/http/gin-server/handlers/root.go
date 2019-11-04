package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/models"
)

func GetRoot(c *gin.Context) {
	data := []models.TextBlock{
		{
			Title: "ようこそう",
			Text:  "You have reached the home page",
		},
	}
	c.JSON(http.StatusOK, data)
}
