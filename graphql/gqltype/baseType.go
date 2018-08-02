package gqltype

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/model"
	"colorme.vn/core/service"
		"github.com/ahl5esoft/golang-underscore"
)

var BaseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Base",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"money": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total revenue of base",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					base := p.Source.(model.Base)

					args := base.Args

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB


					var gen model.Gen

					db.Model(&base).Related(&base.Classes, "base_id")

					classesID := underscore.Pluck(base.Classes, "id").([]uint)

					db.Find(&gen, genID)

					db.Table("registers").
						Select("sum(money) as money").
						Where("? < paid_time AND paid_time < ?", gen.StartTime, gen.EndTime).
						Scopes(model.RegisterByClassesID(classesID)).
						Scopes(model.PaidMoney).Scan(&base)

					return base.Money, nil
				},
			},
		},
	},
)
