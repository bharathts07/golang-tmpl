package blog

import (
	"context"

	"github.com/bharathts07/pokke/internal/database/blog/model"
)

type DB interface {
	// For front page display
	GetBlogSummary(ctx context.Context) ([]model.BlogPreview, error)
	// Main operations
	GetBlog(ctx context.Context, title string) (model.BlogItem, error)
	PostBlog(ctx context.Context, item model.BlogItem) error
	UpdateBlog(ctx context.Context, title string, item model.BlogItem) error
}
