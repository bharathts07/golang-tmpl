package di

import (
	"net/http"

	http2 "github.com/bharathts07/pokke/server/http"
)

func (c *Container) InjectHttpServer(address string) *http.Server {
	if c.Cache.HttpRouter == nil {
		router := http2.CreateRouter(c.InjectDatabase())
		server := &http.Server{
			Addr:    address,
			Handler: router,
		}
		c.Cache.HttpRouter = server
	}
	return c.Cache.HttpRouter
}
