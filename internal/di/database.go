package di

import (
	"github.com/bharathts07/pokke/models/database"
)

func (c *Container) InjectDatabase() database.Client {
	if c.Cache.Database == nil {
		c.InjectLogger().Error("database not yet implemented")
	}
	return c.Cache.Database
}
