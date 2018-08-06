package gqltype

import "github.com/graphql-go/graphql"

var RegisterType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Register",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"money": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
