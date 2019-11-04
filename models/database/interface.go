package database

import (
	"context"
	"github.com/bharathts07/pokke/models/data"
)

// Client specifies the required methods to be implemented
type Client interface {
	GetAll(ctx context.Context) []data.TextBlock
}
