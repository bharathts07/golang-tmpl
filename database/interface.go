package database

import (
	"context"

	"github.com/bharathts07/pokke/internal/models"
)

// Client specifies the required methods to be implemented
type Client interface {
	GetLayout(ctx context.Context, view string) (*models.Layout, error)
	GetComponents(ctx context.Context, view string, components []models.ComponentType) ([]models.Component, error)
}
