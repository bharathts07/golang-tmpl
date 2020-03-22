package di

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/database/blog/mongodb"
	"github.com/bharathts07/pokke/models/database"
	mongopkg "github.com/bharathts07/pokke/pkg/mongodb"
)

func (c *Container) GetDatabase() database.Client {
	if c.cache.database != nil {
		return c.cache.database
	}

	c.GetLogger().Error("database not yet implemented")

	return c.cache.database
}

// GetBlogDB calls GetMongoConnection() and returns the business database implementation blog.DB
func (c *Container) GetBlogDB() blog.DB {
	if c.cache.blogDB != nil {
		return c.cache.blogDB
	}

	db, err := mongodb.New(c.Ctx, c.GetMongoConnection())
	if err != nil {
		c.GetLogger().Error("di.GetBlogDB() :",
			zap.Error(err),
		)
		panic(err)
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
