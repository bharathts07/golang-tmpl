package fire

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"

	"github.com/bharathts07/pokke/internal/database/blog/model"
)

const (
	dbRoot   = "root"
	personal = "personal"
	blogRoot = "blog"

	personalSummary   = "summary"
	blogTypeTechnical = "technical"
)

func (d db) GetBlogSummary(ctx context.Context, limit int) ([]*model.BlogPreview, error) {
	var items = make([]*model.BlogPreview, 0)

	iter := d.client.Collection(dbRoot).Doc(personal).Collection(personalSummary).Limit(limit).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var item model.BlogPreview
		err = doc.DataTo(&item)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (d db) GetBlog(ctx context.Context, title string) (*model.BlogItem, error) {
	var item model.BlogItem

	bSnap, err := d.client.Collection(dbRoot).Doc(personal).Collection(blogRoot).Doc(title).Get(ctx)
	if err != nil {
		return nil, err
	}

	// unmarshal data from database
	err = bSnap.DataTo(&item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (d *db) PostBlog(ctx context.Context, item *model.BlogItem) error {
	_, err := d.client.Collection(dbRoot).Doc(personal).Collection(blogRoot).Doc(item.Title).Set(ctx, item)
	if err != nil {
		return err
	}

	err = UpdatePostSummary(ctx, d.client, item)
	if err != nil {
		return err
	}

	return nil
}

func (d *db) UpdateBlog(ctx context.Context, title string, item *model.BlogItem) error {
	// find the document
	panic("err no t implemented")

	// update the document
}

// UpdatePostSummary constructs a summary from the given item and updates the database with the prepared summary
func UpdatePostSummary(ctx context.Context, client *firestore.Client, item *model.BlogItem) error {
	var firstPhoto string
	if len(item.Photo) > 0 {
		firstPhoto = item.Photo[0]
	}

	// also write a small summary for display in the main page
	sItem := model.BlogPreview{
		ID:      item.ID,
		Title:   item.Title,
		Author:  item.Author,
		Summary: item.Body[:30],
		Tags:    item.Tags,
		Photo:   firstPhoto,
		Created: item.Created,
		Updated: item.Updated,
	}

	_, _, err := client.Collection(dbRoot).Doc(personal).Collection(personalSummary).Add(ctx, sItem)
	if err != nil {
		return err
	}

	return nil
}
