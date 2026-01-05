package middleware

import (
	db "Backend/internal/models"
	"github.com/gin-gonic/gin"
)

const queriesKey = "Queries"

func InjectQueries(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(queriesKey, queries)
		c.Next()
	}
}

func GetQueries(c *gin.Context) *db.Queries {
	value, exists := c.Get(queriesKey)
	if !exists {
		panic("sqlc Queries not found in context")
	}
	queries, ok := value.(*db.Queries)
	if !ok {
		panic("sqlc Queries has incorrect type")
	}
	return queries
}
