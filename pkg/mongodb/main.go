package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewClient returns a mongoDB client after connection and checking the connection
func NewClient(ctx context.Context,
	atlasURL string,
) (*mongo.Client, error) {
	log.Printf("Connecting to mongo db at: %s\n", atlasURL)

	client, err := mongo.NewClient(options.Client().ApplyURI(atlasURL))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("Checking connection to mongo db with a ping\n")

	// check connection with a ping request
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("PING: No error\n")
	return client, nil
}
