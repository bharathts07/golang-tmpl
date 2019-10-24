package di

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/bharathts07/pokke/config"
	"github.com/bharathts07/pokke/database"
)

type Container struct {
	Env *config.Env
	Ctx context.Context

	Cache struct {
		Logger     *zap.Logger
		HttpRouter *http.Server
		Database   database.Client
	}
}

func NewContainer(ctx context.Context, env *config.Env) *Container {
	return &Container{
		Env: env,
		Ctx: ctx,
	}
}
