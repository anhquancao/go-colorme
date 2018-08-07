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
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"rooms": &graphql.Field{
				Type: RoomType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//baseId := p.Source.(model.Base).ID
					//db := service.GetService().DB.DB

					return nil, nil
				},
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

					db.Model(&base).Related(&base.Classes)

					classesID := underscore.Pluck(base.Classes, "id").([]uint)

					db.Find(&gen, genID)

					db.Table("registers").
						Select("sum(money) as money").
						Where("? <= paid_time AND paid_time <= ?", gen.StartTime, gen.EndTime).
						Scopes(model.RegisterByClassesID(classesID)).
						Scopes(model.PaidMoney).Scan(&base)

					return base.Money, nil
				},
			},
			"target_revenue": &graphql.Field{
				Type:        graphql.Int,
				Description: "Total target revenue of base",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					base := p.Source.(model.Base)

					args := base.Args

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var classes []model.Class

					db.Where("base_id = ?", base.ID).Where("gen_id = ?", genID).Find(&classes)

					targetRevenue := model.GetTargetClasses(&classes)

					return targetRevenue, nil
				},
			},
		},
	},
)

var RoomType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Room",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"seat_count": &graphql.Field{
				Type: graphql.Int,
			},
			"avatar_url": &graphql.Field{
				Type: graphql.String,
			},

		},
	},
)
