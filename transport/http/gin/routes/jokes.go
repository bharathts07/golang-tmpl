package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/internal/service/joke"
	"github.com/bharathts07/pokke/transport/http/gin/handlers"
)

func getJokesRouterGroup(v1 *gin.RouterGroup) *gin.RouterGroup {
	jokesGroup := v1.Group("/jokes")
	{
		jokeHandler := handlers.Joke{Service: joke.New()}
		jokesGroup.GET("/", jokeHandler.GetAll)
		jokesGroup.POST("/like/:jokeID", jokeHandler.LikeJoke)
	}

	return jokesGroup
}
