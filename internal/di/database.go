package di

import "github.com/bharathts07/pokke/models/database"

func (c *Container) GetDatabase() database.Client {
	if c.Cache.Database != nil {
		return c.Cache.Database
	}

	c.GetLogger().Error("database not yet implemented")

	return c.Cache.Database
}
