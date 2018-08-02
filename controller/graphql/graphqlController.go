package graphql

import (
	"github.com/gin-gonic/gin"
	"colorme.vn/graphql/util"
	"colorme.vn/graphql/schema"
	"encoding/json"
	"fmt"
)

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"variables"`
}

func GraphQL(c *gin.Context) {
	rawData, _ := c.GetRawData()
	data := string(rawData)
	var request graphQLRequest
	err := json.Unmarshal([]byte(data), &request)
	fmt.Println(err)

	result := util.ExecuteQuery(request.Query, schema.DefaultSchema)
	c.JSON(200, result)
}
