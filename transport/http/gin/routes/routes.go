package routes

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	blogstore "github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/transport/http/gin/handlers"
)

// setupRoutes defines all the rest API endpoints served by this server
func setupRoutes(router *gin.Engine,
	blogdb blogstore.DB,
) {
	router.GET("/", handlers.GetRoot)
	router.Use(static.Serve("/static/", static.LocalFile("./web/app/views/js", true)))
	// api groups is responsible for serving data only
	api := router.Group("/api")
	{
		// version v1 for later deprecation  &/or replacement
		v1 := api.Group("/v1")
		{
			// joke service related api calls
			// ---------------------------------------------------------------------------------------------------
			getJokesRouterGroup(v1)
			// blog service related api calls
			// ---------------------------------------------------------------------------------------------------
			getBlogRouterGroup(v1, blogdb)
		}
	}
}

// CreateRouter creates and configures a server
func CreateRouter(db blogstore.DB) *gin.Engine {
	router := gin.New()
	router.Use(
		// Add MiddleWares here if needed
		cors.New(
			cors.Config{
				AbortOnError:     false,
				AllowAllOrigins:  true,
				AllowedOrigins:   nil,
				AllowOriginFunc:  nil, // don't set this
				AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT"},
				AllowedHeaders:   []string{"*"}, // allow everything
				ExposedHeaders:   []string{"*"}, // allow everything very insecure fix later
				AllowCredentials: true,
				MaxAge:           86400,
			},
		),
	)
	setupRoutes(router, db)
	return router
}
