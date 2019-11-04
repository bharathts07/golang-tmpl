package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/internal/models"
)

func GetRoot(c *gin.Context) {
	data := []models.Component{
		{
			Type: models.Body,
			Data: models.Welcome{
				Message: "ようこそう",
				Home:    "https://0.0.0.0:8080/home",
			},
		},
	}
	c.JSON(http.StatusOK, data)
}
