package gin_server

import (
	"net/http"

	"github.com/bharathts07/pokke/database"
	"github.com/bharathts07/pokke/transport/http/gin-server/routes"
)

// Start creates and starts a gin http server at the address specified
func Start(address string, db database.Client) *http.Server {
	router := routes.CreateRouter(db)

	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
