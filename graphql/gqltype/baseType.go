package gqltype

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/model"
	"colorme.vn/core/service"
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
					baseId := p.Source.(model.Base).ID
					db := service.GetService().DB.DB


					return nil, nil
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
