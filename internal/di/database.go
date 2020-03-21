package di

import (
	"go.uber.org/zap"

	"github.com/bharathts07/pokke/internal/database/blog"
	"github.com/bharathts07/pokke/internal/database/blog/mongodb"
	"github.com/bharathts07/pokke/models/database"
)

func (c *Container) GetDatabase() database.Client {
	if c.Cache.Database != nil {
		return c.Cache.Database
	}

	c.GetLogger().Error("database not yet implemented")

	return c.Cache.Database
}

func (c *Container) GetBlogDatabase() blog.DB {
	if c.Cache.blogDB != nil {
		return c.Cache.blogDB
	}

	con, err := mongodb.New(c.Ctx,
		c.Env.MongoConf.User,
		c.Env.MongoConf.Password,
	)
	if err != nil {
		c.GetLogger().Error("di.GetBlogDatabase() :",
			zap.Error(err),
		)
		panic(err)
	}

	c.Cache.blogDB = con

	return c.Cache.blogDB
}
