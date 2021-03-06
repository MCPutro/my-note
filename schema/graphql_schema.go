package schema

import (
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"time"
)

var (
	userType = graphql.NewObject(graphql.ObjectConfig{
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
		},
	})
	noteType = graphql.NewObject(graphql.ObjectConfig{
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

	respUserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "userResponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: userType,
			},
		},
	})
	respNoteType = graphql.NewObject(graphql.ObjectConfig{
		Name: "noteResponse",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: noteType,
			},
		},
	})

	query    *graphql.Object
	mutation *graphql.Object
)

type GraphQL struct {
	UserService service.UserService
	NoteService service.NoteService
	schema      *handler.Handler
}

func (g *GraphQL) initQueryMutation() {
	query = graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"getAllUser": &graphql.Field{
				Name: "",
				Type: graphql.NewList(userType),
				Args: nil,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user, err := g.UserService.GetAllUser(p.Context)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "Get user list",
			},

			"getUser": &graphql.Field{
				Name: "",
				Type: respUserType,
				Args: graphql.FieldConfigArgument{
					"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					newUser := entity.User{
						Email:    p.Args["email"].(string),
						Password: p.Args["password"].(string),
					}
					if result, err := g.UserService.SignInUser(p.Context, newUser); err != nil {
						return response.Resp{
							Status:  "error",
							Message: err.Error(),
						}, err
					} else {
						return response.Resp{
							Status: "success",
							Data:   result,
						}, nil
					}

				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "get user by email and password",
			},

			"getNoteByUID": &graphql.Field{
				Name: "",
				Type: graphql.NewList(noteType),
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					result, err := g.NoteService.GetNoteByUID(p.Context, p.Args["userId"].(string))
					if err != nil {
						return nil, err
					}
					return result, nil
				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "get all note by userId",
			},
		},
	})

	mutation = graphql.NewObject(graphql.ObjectConfig{
		Name:       "mutation",
		Interfaces: nil,
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Name: "",
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					newUser := entity.User{
						Email:    p.Args["email"].(string),
						Password: p.Args["password"].(string),
					}
					if result, err := g.UserService.CreateNewUser(p.Context, newUser); err != nil {
						return nil, err
					} else {
						return result, nil
					}

				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "create new user",
			},

			"createNote": &graphql.Field{
				Name: "",
				Type: noteType,
				Args: graphql.FieldConfigArgument{
					"userId":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"text":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"themeColor": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var themeColor = p.Args["themeColor"].(string)

					if themeColor == "" {
						themeColor = "theme1"
					}

					newNote := entity.Note{
						Text:       p.Args["text"].(string),
						Visible:    true,
						UserId:     p.Args["userId"].(string),
						ThemeColor: themeColor,
						//User:        entity.User{},
						CreatedDate: time.Now(),
						UpdateDate:  time.Now(),
					}
					note, err := g.NoteService.InsertNewNote(p.Context, newNote)
					if err != nil {
						return nil, err
					}
					return note, nil
				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "create note",
			},

			"removeNote": &graphql.Field{
				Name: "",
				Type: respNoteType,
				Args: graphql.FieldConfigArgument{
					"noteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					err := g.NoteService.Remove(p.Context, p.Args["noteId"].(int))

					if err != nil {
						return response.Resp{
							Status:  "error",
							Message: err.Error(),
						}, err
					}
					return response.Resp{Status: "success"}, nil
				},

				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "remove note to recycle bin (update visible to false)",
			},

			"removeNotePermanent": &graphql.Field{
				Name: "",
				Type: respNoteType,
				Args: graphql.FieldConfigArgument{
					"noteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					err := g.NoteService.RemovePermanent(p.Context, p.Args["noteId"].(int))

					if err != nil {
						return response.Resp{
							Status:  "error",
							Message: err.Error(),
						}, err
					}
					return response.Resp{Status: "success"}, nil
				},

				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "remove note can't be undo",
			},
		},
		IsTypeOf:    nil,
		Description: "",
	})
}

func (g *GraphQL) InitialPath(Route *mux.Router, path string) {
	g.initQueryMutation()

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query,
			Mutation: mutation,
		},
	)
	g.schema = handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	Route.Handle(path, g.schema)
}
