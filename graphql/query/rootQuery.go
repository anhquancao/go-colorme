package query

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/field/base"
	"colorme.vn/graphql/field/analytics"
	"colorme.vn/graphql/field/gen"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"base":            base.FieldBase,
			"bases":           base.FieldBases,
			"gen":             gen.FieldGen,
			"gens":            gen.FieldGens,
			"analytics_sales": analytics.FieldAnalyticsSales,
		},
	},
)
