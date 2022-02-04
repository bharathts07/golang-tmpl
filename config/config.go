package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Env struct {
	// LogLevel is INFO or DEBUG. Default is "INFO".
	LogLevel string `envconfig:"LOG_LEVEL"`
	// Env is used to set GIN log level
	Env string `envconfig:"ENV" default:"development"`
	// HTTPPort is the port at which HTTP endpoints are exposed
	HTTPPort string `envconfig:"HTTP_PORT" default:"5000"`
	// ServiceName used in setting parameters as ID for various monitoring tools
	ServiceName string `envconfig:"SERVICE_NAME" default:"pokke"`
	// Version is a version of the application binary.
	Version string `envconfig:"VERSION"`
	// --------------------------------------------------------------------------------------------
	// BlogDBType holds information regarding the type for database to be used
	BlogDBType string `envconfig:"BLOG_DB_TYPE" default:"mongodb"`
	// Mongo holds config information for connecting to a mongo database
	Mongo struct {
		// URL is the complete url with username and password that can be directly used to connect to mongodb
		URL string `envconfig:"MONGO_URL" default:"mongodb+srv://wippersnapper:1234@cluster0.wvah5.mongodb.net/pokke?retryWrites=true&w=majority" `
	}
	// GCP related env configurations
	GCP struct {
		// ProjectID refers to the GCP project used to deploy the resources required in this code
		ProjectID string `envconfig:"GCP_PROJECT" default:"kouzoh-p-bharath"`
		// CredJSONPath refers to filepath location of the service account credential file
		CredJSONPath string `envconfig:"GCP_CRED_PATH" default:"./keys/kouzoh-p-bharath.json"`
	}
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
	var port string = os.Getenv("PORT")
	var httpPort string = os.Getenv("HTTP_PORT")

	if port == "" {
		port = httpPort
	}

	if port == "" {
		port = "5000"
		_, _ = fmt.Fprint(os.Stdout, "[INFO] No PORT environment variable detected, defaulting to "+port+"\n")
	}
	_, _ = fmt.Fprint(os.Stdout, "[INFO] Selected port is: "+port+"\n")
	return port
}
