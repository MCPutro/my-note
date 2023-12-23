package GraphQL

import (
	"github.com/MCPutro/my-note/GraphQL/schema"
	"github.com/MCPutro/my-note/entity"
	"github.com/MCPutro/my-note/response"
	"github.com/MCPutro/my-note/service"
	"github.com/MCPutro/my-note/util"
	"github.com/graphql-go/graphql"
	graphHandler "github.com/graphql-go/handler"
	"sync"
	"time"
)

var (
	query    *graphql.Object
	mutation *graphql.Object
	once     sync.Once
)

type GraphQL interface {
	InitQueryAndMutation()
	GetHandler() *graphHandler.Handler
}

type graphqlImpl struct {
	userService service.UserService
	noteService service.NoteService
	handlerFunc *graphHandler.Handler
}

func NewGraphQL(userService service.UserService, noteService service.NoteService) GraphQL {
	return &graphqlImpl{userService: userService, noteService: noteService}
}

func (g *graphqlImpl) InitQueryAndMutation() {
	once.Do(func() {
		query = graphql.NewObject(graphql.ObjectConfig{
			Name: "query",
			Fields: graphql.Fields{
				"getAllUser": &graphql.Field{
					Name: "",
					Type: graphql.NewList(schema.User),
					Args: nil,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

						user, err := g.userService.GetAllUser(p.Context)
						if err != nil {
							return nil, err
						}
						return user, nil
					},
					Subscribe:         nil,
					DeprecationReason: "",
					Description:       "Get user list",
				},

				"SignIn": &graphql.Field{
					Name: "",
					Type: schema.User,
					Args: graphql.FieldConfigArgument{
						"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
						"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						newUser := entity.User{
							Email:    p.Args["email"].(string),
							Password: p.Args["password"].(string),
						}
						if result, err := g.userService.SignInUser(p.Context, newUser); err != nil {
							return nil, err
						} else {
							return result, nil
						}

					},
					Subscribe:         nil,
					DeprecationReason: "",
					Description:       "get user by email and password",
				},

				"getNoteByUID": &graphql.Field{
					Name: "",
					Type: graphql.NewList(schema.Note),
					Args: graphql.FieldConfigArgument{
						"userId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

						result, err := g.noteService.GetNoteByUID(p.Context, p.Args["userId"].(string))
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
				"SignUp": &graphql.Field{
					Name: "",
					Type: schema.User,
					Args: graphql.FieldConfigArgument{
						"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
						"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

						newUser := entity.User{
							Email:    p.Args["email"].(string),
							Password: p.Args["password"].(string),
						}
						if result, err := g.userService.CreateNewUser(p.Context, newUser); err != nil {
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
					Type: schema.Note,
					Args: graphql.FieldConfigArgument{
						"userId":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
						"text":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
						"themeColor": &graphql.ArgumentConfig{Type: graphql.String},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

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
						note, err := g.noteService.InsertNewNote(p.Context, newNote)
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
					Type: schema.Note,
					Args: graphql.FieldConfigArgument{
						"noteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

						err := g.noteService.Remove(p.Context, p.Args["noteId"].(int))

						if err != nil {
							return nil, err
						}
						return response.Resp{Status: "success"}, nil
					},

					Subscribe:         nil,
					DeprecationReason: "",
					Description:       "remove note to recycle bin (update visible to false)",
				},

				"removeNotePermanent": &graphql.Field{
					Name: "",
					Type: schema.Note,
					Args: graphql.FieldConfigArgument{
						"noteId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						auth := p.Context.Value("Authorization").(string)
						if _, err2 := util.ValidateToken(auth); err2 != nil {
							return nil, err2
						}

						err := g.noteService.RemovePermanent(p.Context, p.Args["noteId"].(int))

						if err != nil {
							return nil, err
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

		newSchema, _ := graphql.NewSchema(
			graphql.SchemaConfig{
				Query:    query,
				Mutation: mutation,
			})

		newHandler := graphHandler.New(&graphHandler.Config{
			Schema: &newSchema,
		})

		g.handlerFunc = newHandler
	})
}

func (g *graphqlImpl) GetHandler() *graphHandler.Handler {
	return g.handlerFunc
}
