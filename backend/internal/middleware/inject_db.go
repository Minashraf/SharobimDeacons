package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

const dbKey = "Database"

func InjectDatabase(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(dbKey, db)
		c.Next()
	}
}

func GetDatabase(c *gin.Context) *sql.DB {
	value, exists := c.Get(dbKey)
	if !exists {
		panic("Database instance not found in context")
	}
	queries, ok := value.(*sql.DB)
	if !ok {
		panic("Database has incorrect type")
	}
	return queries
}
