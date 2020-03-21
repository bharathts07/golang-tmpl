package di

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/models/database"
)

// Container is a DI container with all the dependencies listed and methods associated with this struct to fetch those
// dependencies
type Container struct {
	Env *config.Env
	Ctx context.Context

	Cache struct {
		Logger     *zap.Logger
		HTTPServer *http.Server
		Database   database.Client
		// database for blog management
		blogDB blog.DB
	}
}

func NewContainer(ctx context.Context, env *config.Env) *Container {
	return &Container{
		Env: env,
		Ctx: ctx,
	}
}
