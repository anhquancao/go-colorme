package base

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/gqltype"
	"colorme.vn/model"
	"colorme.vn/core/service"
)

var FieldBase = &graphql.Field{
	Type:        gqltype.BaseType,
	Description: "Get base by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"gen_id": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Get info base by gen id",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		id, ok := p.Args["id"].(int)

		var base model.Base

		if ok {
			db.Find(&base, id)
		}

		base.Args = p.Args
		return base, nil
	},
}

var FieldBases = &graphql.Field{
	Type:        graphql.NewList(gqltype.BaseType),
	Description: "Get bases list",
	Args: graphql.FieldConfigArgument{
		"gen_id": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "Get info base by gen id",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var bases []model.Base
		db := service.GetService().DB.DB
		db.Order("created_at asc").Find(&bases)

		for _, base := range bases {
			base.Args = p.Args
		}

		return bases, nil
	},
}


