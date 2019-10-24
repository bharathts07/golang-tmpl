package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	// LogLevel is INFO or DEBUG. Default is "INFO".
	LogLevel string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	// Env is used to set GIN log level
	Env string `envconfig:"ENV" required:"true"`
	// HTTPPort is the port at which HTTP endpoints are exposed
	HTTPPort string `envconfig:"PORT" default:"5000"`
	// ServiceName used in setting parameters as ID for various monitoring tools
	ServiceName string `envconfig:"SERVICE_NAME" default:"pokke"`
	// Version is a version of the application binary.
	Version string
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	_, _ = fmt.Fprint(os.Stdout, "[DEBUG] Beginning to parse ENV ")
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}
	_, _ = fmt.Fprint(os.Stdout, "[DEBUG] Set environment values are:\n ")
	for _, keyVal := range os.Environ() {
		fmt.Println(keyVal)
	}
	return &env, nil
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return port
}
