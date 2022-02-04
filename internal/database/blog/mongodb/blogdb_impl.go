package mongodb

import (
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bharathts07/pokke/internal/database/blog/model"
)

func (d db) GetBlogSummary(ctx context.Context, limit int) ([]*model.BlogPreview, error) {
	_ = d.client.Database("blog").Collection("blog")
	return nil, nil
}

func (d db) GetBlog(ctx context.Context, title string) (*model.BlogItem, error) {
	collection := d.client.Database("blog").Collection("blog")
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := collection.FindOne(ctx, bson.D{{"title", title}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}

	return nil, nil
}

func (d db) PostBlog(ctx context.Context, item *model.BlogItem) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := d.client.Database("blog").Collection("blog")
	res, err := collection.InsertOne(ctx, bson.D{{item.Title, item}})
	if err != nil {
		return err
	}
	fmt.Printf(fmt.Sprint(res.InsertedID))

	return nil
}

func (d db) UpdateBlog(ctx context.Context, title string, item *model.BlogItem) error {
	_ = d.client.Database("blog").Collection("blog")
	return nil
}
