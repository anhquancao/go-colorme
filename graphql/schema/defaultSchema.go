package schema

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/query"
)

var DefaultSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:query.RootQuery,
	},
)