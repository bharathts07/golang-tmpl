package http

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/database"
	"github.com/bharathts07/pokke/server/middleware"
)

// CreateRouter creates and configures a server
func CreateRouter(db database.Client) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.DB(db))
	setupRoutes(router)
	return router
}
