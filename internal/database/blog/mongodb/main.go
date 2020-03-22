package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bharathts07/pokke/internal/database/blog"
)

// db provides concrete implementation for blog.DB interface
type db struct {
	client *mongo.Client
}

// New returns the concrete implementation for blog.DB
func New(ctx context.Context,
	client *mongo.Client,
) (blog.DB, error) {
	return &db{
		client: client,
	}, nil
}
