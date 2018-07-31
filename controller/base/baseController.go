package base

import (
	"github.com/gin-gonic/gin"
	"colorme.vn/graphql/util"
	"colorme.vn/graphql/schema"
)

func BaseGraphQL(c *gin.Context) {
	query := c.Query("query")
	result := util.ExecuteQuery(query, schema.BaseSchema)
	c.JSON(200, result)
}