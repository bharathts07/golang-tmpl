package di

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/server/http"
)

func (c *Container) InjectHttpGinRouter() *gin.Engine {
	if c.Cache.HttpRouter == nil {
		router := http.CreateRouter(c.Cache.Database)
		c.Cache.HttpRouter = router
	}
	return c.Cache.HttpRouter
}
