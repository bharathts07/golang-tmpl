package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, projectID, cred string) (*firestore.Client, error) {
	client, err := firestore.NewClient(
		ctx,
		projectID,
		option.WithCredentialsFile(cred),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
