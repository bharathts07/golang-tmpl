package di

import (
	"net/http"

	"github.com/bharathts07/pokke/transport/http/gin"
)

func (c *Container) GetHTTPServer(address string) *http.Server {
	if c.Cache.HTTPRouter == nil {
		server := gin.Start(address)
		c.Cache.HTTPRouter = server
	}
	return c.Cache.HTTPRouter
}
