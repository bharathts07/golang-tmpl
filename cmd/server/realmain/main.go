package realmain

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/di"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

// Execute contains the implementation and logic for the server. It returns an exit code indicating exit status
func Execute() int {
	_, _ = fmt.Fprint(os.Stdout, "[INFO] Beginning Execution\n")
	// Context for server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Parse environment variables here
	//env, _ := config.ReadFromEnv()
	_, _ = fmt.Fprint(os.Stdout, "[INFO] Creating container\n")
	// Create a DI container that simplifies
	container := di.NewContainer(ctx, nil)
	// address is the address at which the server listens
	address := fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort())
	// Create http server with the required configurations
	httpServer := container.InjectHttpServer(address)

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		_, _ = fmt.Fprint(os.Stdout, "[INFO] Starting server at address\n", address,"\n")
		return httpServer.ListenAndServe()
	})
	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	// Waiting for SIGTERM or Interrupt signal. If server receives them,
	// gRPC server and http server will shutdown gracefully.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
		_, _ = fmt.Fprint(os.Stdout, "\n[INFO] Received SIGTERM, exiting server gracefully\n")
	case <-ctx.Done():
	}
	// Remember to close all connections like db here
	return exitOK
}
