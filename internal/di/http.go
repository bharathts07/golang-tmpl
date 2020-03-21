package di

import (
	"fmt"
	"net/http"

	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/transport/http/gin"
)

// GetHTTPServer returns a gin implementation for httpServer
func (c *Container) GetHTTPServer() *http.Server {
	if c.Cache.HTTPServer != nil {
		return c.Cache.HTTPServer
	}

	address := fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort())
	server := gin.Start(address)

	c.Cache.HTTPServer = server

	return c.Cache.HTTPServer
}
