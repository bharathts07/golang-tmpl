package blog

import (
	"context"
	"fmt"
	"time"

	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/database/blog/model"
	"github.com/bharathts07/pokke/internal/service/blog/utils"
)

type impl struct {
	db blog.DB
}

// NewService take the required database dependency and return the implementation for the blog service
func NewService(db blog.DB) Service {
	return &impl{db: db}
}

func (svc impl) ReadBlog(ctx context.Context, title string) (*model.BlogItem, error) {
	blogItem, err := svc.db.GetBlog(ctx, title)
	if err != nil {
		return nil, fmt.Errorf("svc.ReadBlog() : %+v", err)
	}

	return blogItem, nil
}

func (svc impl) SaveBlog(ctx context.Context, blogItem *model.BlogItem) error {
	blogItem.Author = "self"
	blogItem.Created = time.Now()
	blogItem.Updated = time.Now()
	blogItem.ID = utils.GetUniqueID()

	err := svc.db.PostBlog(ctx, blogItem)
	if err != nil {
		return fmt.Errorf("svc.SaveBlog() : %+v", err)
	}

	return nil
}

func (svc impl) UpdateBlog(ctx context.Context, title string, b *model.BlogItem) error {
	panic("implement me")
}

func (svc impl) BlogPreview(ctx context.Context, limit int) ([]*model.BlogPreview, error) {
	val, err := svc.db.GetBlogSummary(ctx, limit)
	if err != nil {
		return nil, fmt.Errorf("svc.BlogPreview() : %+v", err)
	}

	return val, nil
}
