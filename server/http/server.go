package http

import (
	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/database"
	"strconv"
)

// StartServer starts given server, supporting graceful shutdown of the server
func StartServer(db database.Client, env *config.Env) {
	router := CreateRouter(db)
	serverPort := "0.0.0.0:"+strconv.Itoa(env.GRPCPort)
	_ = router.Run(serverPort)
}
