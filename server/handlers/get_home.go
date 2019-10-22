package handlers

import (
	"github.com/bharathts07/pokke/internal/models"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	data := models.Home{
		Layout:models.HomeLayout{
			Alignment:    "Center",
			Order:        models.OrderMap{},
			ComponentIDs: nil,
		},
	}
	c.JSON(200, data)
}
