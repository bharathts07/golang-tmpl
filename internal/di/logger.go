package di

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/bharathts07/pokke/pkg/log"
)

func (c *Container) GetLogger() *zap.Logger {
	if c.Cache.Logger == nil {
		logger, err := log.New("Debug")
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, "[ERROR] Initializing logger")
			panic(err)
		}
		c.Cache.Logger = logger
	}
	return c.Cache.Logger
}
