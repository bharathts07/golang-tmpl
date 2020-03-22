package di

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/transport/http/gin"
)

// GetHTTPServer returns a gin implementation for httpServer
func (c *Container) GetHTTPServer() *http.Server {
	if c.cache.hTTPServer != nil {
		return c.cache.hTTPServer
	}

	address := fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort())
	log.Printf("[INFO] Starting http server at address : %s", address)

	db := c.GetBlogDB()

	server := gin.Start(address, db)

	c.cache.hTTPServer = server

	return c.cache.hTTPServer
}
