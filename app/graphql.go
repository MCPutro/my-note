package app

import (
	"github.com/MCPutro/my-note/schema"
	"github.com/MCPutro/my-note/service"
)

func NewGraphQL(userService service.UserService, noteService service.NoteService) *schema.GraphQL {
	return &schema.GraphQL{
		UserService: userService,
		NoteService: noteService,
	}
}
