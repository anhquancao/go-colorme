package query

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/field/base"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"base":  base.FieldBase,
			"bases": base.FieldBases,
		},
	},
)
