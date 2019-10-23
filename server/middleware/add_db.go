package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/bharathts07/pokke/database"
	"github.com/bharathts07/pokke/global"
)

//DB middleware attaches a database connection to gin's Context
func DB(db database.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(global.ContextDBName, db)
		c.Next()
	}
}

// GetDB extracts the stored database client from context
func GetDB(c *gin.Context) database.Client {
	return c.MustGet(global.ContextDBName).(database.Client)
}
