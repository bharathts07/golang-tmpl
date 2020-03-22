package gin

import (
	"net/http"

	"github.com/bharathts07/pokke/transport/http/gin/routes"
)

// Start creates and starts a gin http server at the address specified
func Start(address string) *http.Server {
	router := routes.CreateRouter()

	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
