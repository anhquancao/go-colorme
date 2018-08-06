package analytics

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/gqltype"
)

var FieldAnalyticsSales = &graphql.Field{
	Type:        gqltype.AnalyticSalesType,
	Description: "Get analytic sales",
	Args: graphql.FieldConfigArgument{
		"gen_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return p.Args, nil
	},
}
