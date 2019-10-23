package http

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/server/handlers"
)

// setupRoutes defines all the rest API endpoints served by this server
func setupRoutes(router *gin.Engine) {

	router.GET("/", handlers.GetWelcomeComponents)
	//v1 := router.Group("/v1",)
	// Group all the resources for displaying home page
	//home := v1.Group("/home")
	//home.GET("/",home)
	//home.GET("/layout")
	//home.POST("/components")

}
