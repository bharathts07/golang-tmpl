package di

import (
	"net/http"

	"github.com/bharathts07/pokke/transport/http/gin-server"
)

func (c *Container) InjectHttpServer(address string) *http.Server {
	if c.Cache.HttpRouter == nil {
		server := gin_server.Start(address, c.InjectDatabase())
		c.Cache.HttpRouter = server
	}
	return c.Cache.HttpRouter
}
