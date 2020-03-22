package mongodb

import (
	"context"

	"github.com/bharathts07/pokke/internal/database/blog/model"
)

func (d db) GetBlogSummary(ctx context.Context,limit int) ([]*model.BlogPreview, error) {
	panic("implement me")
}

func (d db) GetBlog(ctx context.Context, title string) (*model.BlogItem, error) {
	panic("implement me")
}

func (d db) PostBlog(ctx context.Context, item *model.BlogItem) error {
	panic("implement me")
}

func (d db) UpdateBlog(ctx context.Context, title string, item *model.BlogItem) error {
	panic("implement me")
}
