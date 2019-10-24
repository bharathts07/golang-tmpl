package http

import (
	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/database"
	log2 "github.com/bharathts07/pokke/internal/pkg/log"
	"go.uber.org/zap"
	"strconv"
)

// StartServer starts given server, supporting graceful shutdown of the server
func StartServer(db database.Client, env *config.Env) {
	router := CreateRouter(db)
	//Start logger for logging received env values
	log,_ := log2.New("Debug")
	log.Info("Parsed ENV values",
		zap.Any("Values",env),
		)

	serverPort := "0.0.0.0:"+strconv.Itoa(env.GRPCPort)
	_ = router.Run(serverPort)
}
