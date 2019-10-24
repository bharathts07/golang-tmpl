package main

import (
	"fmt"
	"os"

	"github.com/bharathts07/pokke/cmd/server/realmain"
)

// version is git commit at the time of build of this service binary. version is overwritten by -ldflags
// while building.
var version string

// serviceName is service name of this microservices. serviceName
// is overwritten using -ldflags while building.
var serviceName string

func main() {
	_, _ = fmt.Fprintf(os.Stdout, "[INFO] Starting service\n")
	os.Exit(realmain.Execute())
}
