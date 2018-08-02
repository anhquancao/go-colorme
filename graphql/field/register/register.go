package register

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/graphql/gqltype"
	"colorme.vn/model"
	"colorme.vn/core/service"
	)

var FieldRegister = &graphql.Field{
	Type:        gqltype.RegisterType,
	Description: "Get register by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		id, ok := p.Args["id"].(int)

		var register model.Register

		if ok {
			db.Find(&register, id)
		}
		return register, nil
	},
}
