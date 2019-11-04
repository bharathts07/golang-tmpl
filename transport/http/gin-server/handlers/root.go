package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/models"
)

// GetRoot prints a json content when the root address is clicked
// TODO : To be fixed in the future, to display a web page
func GetRoot(c *gin.Context) {
	data := []models.TextBlock{
		{
			Title: "ようこそう",
			Text:  "You have reached the home page",
		},
	}
	c.JSON(http.StatusOK, data)
}
