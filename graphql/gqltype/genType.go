package gqltype

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/model"
)

var GenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Gen",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"start_time": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					gen := p.Source.(model.Gen)
					return gen.StartTime.Unix(), nil
				},
			},
			"end_time": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					gen := p.Source.(model.Gen)
					return gen.EndTime.Unix(), nil
				},
			},
		},
	},
)
