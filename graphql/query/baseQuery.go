package query

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/gqltype"
	"colorme.vn/model"
)

var bases = [3]model.Base{
	{ID: 1, Name: "Hanoi", Address: "175 chua lang"},
	{ID: 2, Name: "Sai Gon", Address: "Quan 1"},
	{ID: 3, Name: "Da Nang", Address: "Cau truong tien"},
}

var BaseQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"base": &graphql.Field{
				Type:        gqltype.BaseType,
				Description: "Get base by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, base := range bases {
							if int(base.ID) == id {
								return base, nil
							}
						}
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(gqltype.BaseType),
				Description: "Get bases list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return bases, nil
				},
			},
		},
	},
)
