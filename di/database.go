package di

import (
	"github.com/bharathts07/pokke/database"
	"github.com/bharathts07/pokke/database/fakedb"
)

func (c *Container) InjectDatabase() database.Client {
	if c.Cache.Database == nil {
		db := fakedb.New()
		c.Cache.Database = db
	}
	return c.Cache.Database
}
