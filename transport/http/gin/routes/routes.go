package routes

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	blogstore "github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/service/blog"
	"github.com/bharathts07/pokke/internal/service/joke"
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
			jokeHandler := handlers.Joke{Service: joke.New()}
			v1.GET("/jokes", jokeHandler.GetAll)
			v1.POST("/jokes/like/:jokeID", jokeHandler.LikeJoke)

			// blog service related api calls
			// ---------------------------------------------------------------------------------------------------
			blogGroup := v1.Group("/blog")
			{
				blogHandler := handlers.Blog{Service: blog.NewService(blogdb)}
				blogGroup.GET("/personal", blogHandler.GetAll)
				blogGroup.GET("/personal/:title", blogHandler.GetBlog)
				blogGroup.POST("/personal/:title", blogHandler.PostBlog)
			}
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
