package realmain

import (
	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/database/fakedb"
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
	// get reference to a database from which to retrieve the data
	db := fakedb.New()
	env,_ := config.ReadFromEnv()
	// start the http server to begin serving backend requests
	http.StartServer(db,env)
	return exitOK
}
