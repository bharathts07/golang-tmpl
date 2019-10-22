package realmain

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/server/http"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

// Execute contains the implementation and logic for the server. It returns an exit code indicating exit status
func Execute(_ []string) int {
	r := gin.Default()
	http.StartServer(r)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		return exitError
	}
	return exitOK
}
