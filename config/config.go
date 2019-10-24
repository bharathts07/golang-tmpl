package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
)

type Env struct {
	// LogLevel is INFO or DEBUG. Default is "INFO".
	LogLevel string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	// Env is used to set GIN log level
	Env string `envconfig:"ENV" required:"true"`
	// HTTPPort is the port at which HTTP endpoints are exposed
	HTTPPort int `envconfig:"PORT" default:"5000"`
	// ServiceName used in setting parameters as ID for various monitoring tools
	ServiceName string `envconfig:"SERVICE_NAME" default:"pokke"`
	// Version is a version of the application binary.
	Version string
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	if err := env.validate(); err != nil {
		return nil, errors.Wrap(err, "validation failed")
	}

	return &env, nil
}

// validate validates
func (e *Env) validate() error {
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{
			e.Env != envDevelopment && e.Env != envProduction,
			fmt.Sprintf("invalid env is specified: %q", e.Env),
		},

		// Add your own validation here
	}

	for _, check := range checks {
		if check.bad {
			return errors.Errorf(check.errMsg)
		}
	}

	return nil
}
