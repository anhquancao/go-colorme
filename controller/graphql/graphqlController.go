package graphql

import (
	"github.com/gin-gonic/gin"
	"colorme.vn/graphql/util"
	"colorme.vn/graphql/schema"
)

func GraphQL(c *gin.Context) {
	query := c.Query("query")
	result := util.ExecuteQuery(query, schema.DefaultSchema)
	c.JSON(200, result)
}
