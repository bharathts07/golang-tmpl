package di

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"go.mongodb.org/mongo-driver/mongo"
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

	cache struct {
		logger     *zap.Logger
		hTTPServer *http.Server
		// database for blog management
		mongoDB    *mongo.Client
		fireClient *firestore.Client

		// business specific dependencies constructed locally
		database database.Client
		blogDB   blog.DB
	}
}

func NewContainer(ctx context.Context, env *config.Env) *Container {
	return &Container{
		Env: env,
		Ctx: ctx,
	}
}
