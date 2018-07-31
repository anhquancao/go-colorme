package schema

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/query"
)

var BaseSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:query.BaseQuery,
	},
)