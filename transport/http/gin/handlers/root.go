package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/models/data"
)

// GetRoot prints a json content when the root address is clicked
func GetRoot(c *gin.Context) {
	d := []data.TextBlock{
		{
			Title: "ようこそう",
			Text:  "You have reached the home page",
		},
	}
	c.JSON(http.StatusOK, d)
}
