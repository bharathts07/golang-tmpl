package http

import "github.com/gin-gonic/gin"

// CreateRouter creates and configures a server
func CreateRouter() *gin.Engine {
	router := gin.Default()
	setupRoutes(router)
	return router
}
