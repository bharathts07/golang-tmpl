package fire

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/bharathts07/pokke/internal/database/blog"
)

// db provides concrete implementation for blog.DB interface
type db struct {
	client *firestore.Client
}

// New returns the concrete implementation for blog.DB
func New(ctx context.Context,
	client *firestore.Client,
) (blog.DB, error) {
	return &db{
		client: client,
	}, nil
}
