package blog

import (
	"context"

	"github.com/bharathts07/pokke/internal/database/blog/model"
)

type Service interface {
	ReadBlog(ctx context.Context, title string) (*model.BlogItem, error)
	SaveBlog(ctx context.Context, blog *model.BlogItem) error
	UpdateBlog(ctx context.Context, title string, blog *model.BlogItem) error
	BlogPreview(ctx context.Context, limit int) ([]*model.BlogPreview, error)
}
