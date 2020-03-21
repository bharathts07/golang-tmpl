package handlers

import (
	data2 "github.com/bharathts07/pokke/models/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoot prints a json content when the root address is clicked
// TODO : To be fixed in the future, to display a web page
func GetRoot(c *gin.Context) {
	data := []data2.TextBlock{
		{
			Title: "ようこそう",
			Text:  "You have reached the home page",
		},
	}
	c.JSON(http.StatusOK, data)
}
