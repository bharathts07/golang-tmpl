package routes

import (
	"github.com/gin-gonic/gin"

	blog2 "github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/service/blog"
	"github.com/bharathts07/pokke/transport/http/gin/handlers"
)

func getBlogRouterGroup(v1 *gin.RouterGroup, db blog2.DB) *gin.RouterGroup {
	blogGroup := v1.Group("/blog")
	{
		blogHandler := handlers.Blog{Service: blog.NewService(db)}
		blogGroup.GET("/personal/:title", blogHandler.GetBlog)
		blogGroup.POST("/personal/:title", blogHandler.PostBlog)
		blogGroup.GET("/all/:limit", blogHandler.GetAll)
	}
	return blogGroup
}
