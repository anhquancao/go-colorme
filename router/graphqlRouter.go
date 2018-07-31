package router

import (
	"colorme.vn/core"
	"colorme.vn/controller/graphql"
)

func RegisterGraphQLRouter(context *core.Context) {
	server := context.Server
	server.POST("graphql", graphql.GraphQL)
	server.GET("graphql", graphql.GraphQL)
}
