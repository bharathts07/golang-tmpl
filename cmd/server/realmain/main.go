package realmain

import (
	"context"
	"fmt"
	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/di"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

const (
	// exit is exit code which is returned by realMain function.
	// exit code is passed to os.Exit function.
	exitOK = iota
	exitError
)

// Execute contains the implementation and logic for the server. It returns an exit code indicating exit status
func Execute(_ []string) int {
	// Context for server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Parse environment variables here
	env , _ := config.ReadFromEnv()
	// Create a DI container that simplifies
	container := di.NewContainer(ctx,env)
	// address is the address at which the server listens
	address := fmt.Sprintf("%s:%s","0.0.0.0",env.HTTPPort)
	// Create http server with the required configurations
	httpServer := container.InjectHttpServer(address)

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
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
		_, _ = fmt.Fprint(os.Stdout, "[INFO] Received SIGTERM, exiting server gracefully")
	case <-ctx.Done():
	}
	if err := wg.Wait(); err != nil {
		_, _ = fmt.Fprint(os.Stdout, "[INFO] unhandled error received")
		_, _ = fmt.Fprint(os.Stdout, err)
		return exitError
	}
	return exitOK
}
