package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/pkg/mongodb"
)

// db provides concrete implementation for blog.DB interface
type db struct {
	client *mongo.Client
}

// New returns the concrete implementation for blog.DB
func New(ctx context.Context,
	user, pass string,
) (blog.DB, error) {
	mClient, err := mongodb.NewClient(ctx, user, pass)
	if err != nil {
		return nil, fmt.Errorf("mongodb.New() : %+v", err)
	}

	return &db{
		client: mClient,
	}, nil
}
