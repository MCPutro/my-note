package schema

import (
	"github.com/graphql-go/graphql"
)

var (
	User = graphql.NewObject(graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	Note = graphql.NewObject(graphql.ObjectConfig{
		Name: "note",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"themeColor": &graphql.Field{
				Type: graphql.String,
			},
			"createdDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"updateDate": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

	RespUser = graphql.NewObject(graphql.ObjectConfig{
		Name: "userResponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: User,
			},
		},
	})

	RespNote = graphql.NewObject(graphql.ObjectConfig{
		Name: "noteResponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: Note,
			},
		},
	})
)
