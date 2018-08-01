package gqltype

import "github.com/graphql-go/graphql"

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
		},
	},
)
