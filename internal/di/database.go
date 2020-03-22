package di

import (
	"cloud.google.com/go/firestore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/database/blog/fire"
	"github.com/bharathts07/pokke/internal/database/blog/mongodb"
	"github.com/bharathts07/pokke/models/database"
	firepkg "github.com/bharathts07/pokke/pkg/firestore"
	mongopkg "github.com/bharathts07/pokke/pkg/mongodb"
)

func (c *Container) GetDatabase() database.Client {
	if c.cache.database != nil {
		return c.cache.database
	}

	c.GetLogger().Error("database not yet implemented")

	return c.cache.database
}

// GetFireStoreClient connects to the firestore database and returns the client
func (c *Container) GetFireStoreClient() *firestore.Client {
	if c.cache.fireClient != nil {
		return c.cache.fireClient
	}

	client, err := firepkg.NewClient(c.Ctx, c.Env.GCP.ProjectID, c.Env.GCP.CredJSONPath)
	if err != nil {
		c.GetLogger().Error("connect to firestore",
			zap.Error(err),
		)
		panic(err)
	}

	c.GetLogger().Info("connected to firestore")
	c.cache.fireClient = client

	return c.cache.fireClient
}

// GetBlogDB calls GetMongoConnection() and returns the business database implementation blog.DB
func (c *Container) GetBlogDB() blog.DB {
	if c.cache.blogDB != nil {
		return c.cache.blogDB
	}
	var db blog.DB
	var err error

	switch c.Env.BlogDBType {
	case "firestore":
		db, err = fire.New(c.Ctx, c.GetFireStoreClient())
		if err != nil {
			c.GetLogger().Error("di.GetBlogDB() :",
				zap.Error(err),
			)
			panic(err)
		}
	case "mongodb":
		db, err = mongodb.New(c.Ctx, c.GetMongoConnection())
		if err != nil {
			c.GetLogger().Error("di.GetBlogDB() :",
				zap.Error(err),
			)
			panic(err)
		}
	}

	c.cache.blogDB = db

	return c.cache.blogDB
}

// GetMongoConnection connects to the mongo database using the config user and password
func (c *Container) GetMongoConnection() *mongo.Client {
	if c.cache.blogDB != nil {
		return c.cache.mongoDB
	}

	con, err := mongopkg.NewClient(c.Ctx, c.Env.Mongo.URL)
	if err != nil {
		c.GetLogger().Error("di.GetMongoConnection() :",
			zap.Error(err),
		)
		panic(err)
	}

	c.cache.mongoDB = con

	return c.cache.mongoDB
}
