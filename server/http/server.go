package http

// StartServer starts given server, supporting graceful shutdown of the server
func StartServer() {
	router := CreateRouter()
	_ := router.Run(":8080")
}