package http

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/server/handlers"
)

// setupRoutes defines all the rest API endpoints served by this server
func setupRoutes(router *gin.Engine) {

	router.GET("/", handlers.GetWelcomeComponents)
	router.Use(static.Serve("/static/",static.LocalFile("./assets/views/js",true)))
	// api groups is responsible for serving data only
	api := router.Group("/api",)
	{
		// version v1 for later deprecation  &/or replacement
		v1 := api.Group("/v1")
		{
			v1.GET("/jokes", handlers.JokeHandler )
			v1.POST("/jokes/like/:jokeID",handlers.LikeJoke)
		}
	}

}
