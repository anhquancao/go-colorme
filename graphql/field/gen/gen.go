package gen

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/gqltype"
	"colorme.vn/model"
	"colorme.vn/core/service"
	)

var FieldGen = &graphql.Field{
	Type:        gqltype.GenType,
	Description: "Get base by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		genID, ok := p.Args["id"].(int)
		if !ok {
			genID = int(model.GetCurrentGen().ID)
		}

		var gen model.Gen

		db.Find(&gen, genID)

		gen.Args = p.Args
		return gen, nil
	},
}

var FieldGens = &graphql.Field{
	Type:        graphql.NewList(gqltype.GenType),
	Description: "Get gens list",
	Args: graphql.FieldConfigArgument{
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var gens []model.Gen
		db := service.GetService().DB.DB
		db.Order("created_at asc").Find(&gens)

		for _, base := range gens {
			base.Args = p.Args
		}

		return gens, nil
	},
}
