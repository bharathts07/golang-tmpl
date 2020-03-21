package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	AtlasURL = "mongodb+srv://%s:%s@cluster0-s4ipk.mongodb.net/test?retryWrites=true&w=majority"
)

// NewClient returns a mongoDB client after connection and checking the connection
func NewClient(ctx context.Context,
	user, pass string,
) (*mongo.Client, error) {
	connectionString := fmt.Sprintf(AtlasURL, user, pass)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// check connection with a ping request
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
