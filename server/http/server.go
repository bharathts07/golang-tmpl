package http

import "github.com/bharathts07/pokke/database"

// StartServer starts given server, supporting graceful shutdown of the server
func StartServer(db database.Client) {
	router := CreateRouter(db)
	_ = router.Run(":8080")
}
