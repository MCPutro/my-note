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
		Name: "User",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.String,
			},
			"Email": &graphql.Field{
				Type: graphql.String,
			},
			"Password": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	noteType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Note",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"Text": &graphql.Field{
				Type: graphql.String,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UserID": &graphql.Field{
				Type: graphql.String,
			},
			"ThemeColor": &graphql.Field{
				Type: graphql.String,
			},
			"CreatedDate": &graphql.Field{
				Type: graphql.DateTime,
			},
			"UpdateDate": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	})

	respUserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserResponse",
		Fields: graphql.Fields{
			"Status": &graphql.Field{
				Type: graphql.String,
			},
			"Message": &graphql.Field{
				Type: graphql.String,
			},
			"Data": &graphql.Field{
				Type: userType,
			},
		},
	})
	respNoteType = graphql.NewObject(graphql.ObjectConfig{
		Name: "NoteResponse",
		Fields: graphql.Fields{
			"Status": &graphql.Field{
				Type: graphql.String,
			},
			"Message": &graphql.Field{
				Type: graphql.String,
			},
			"Data": &graphql.Field{
				Type: noteType,
			},
		},
	})

	query    *graphql.Object
	mutation *graphql.Object
)

type GraphQL struct {
	UserService *service.UserService
	NoteService *service.NoteService
	schema      *handler.Handler
	Route       *mux.Router
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
					user, err := g.UserService.GetAllUser()
					if err != nil {
						return nil, err
					}
					return user, nil
				},
				Subscribe:         nil,
				DeprecationReason: "",
				Description:       "Get book list",
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
					if result, err := g.UserService.SignInUser(newUser); err != nil {
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
				Description:       "",
			},

			"getNoteByUID": &graphql.Field{
				Name: "",
				Type: graphql.NewList(noteType),
				Args: graphql.FieldConfigArgument{
					"UserId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					result, err := g.NoteService.GetNoteByUID(p.Args["UserId"].(string))
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
					if result, err := g.UserService.CreateNewUser(newUser); err != nil {
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
					note, err := g.NoteService.InsertNewNote(newNote)
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
					"NoteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					err := g.NoteService.Remove(p.Args["NoteId"].(int))

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
				Description:       "remove note to recycle bin",
			},

			"removeNotePermanent": &graphql.Field{
				Name: "",
				Type: respNoteType,
				Args: graphql.FieldConfigArgument{
					"NoteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					err := g.NoteService.RemovePermanent(p.Args["NoteId"].(int))

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

func (g *GraphQL) InitialPath(path string) {
	g.initQueryMutation()

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query,
			Mutation: mutation,
		},
	)
	g.schema = handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		//Playground: true,
	})

	g.Route.Handle(path, g.schema)
}
