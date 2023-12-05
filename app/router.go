package app

import (
	"fmt"
	myGraphQL "github.com/MCPutro/my-note/GraphQL"
	myGraphQLHandler "github.com/MCPutro/my-note/GraphQL/handler"
	"github.com/MCPutro/my-note/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userController controller.UserController, noteController controller.NoteController, graphql myGraphQL.GraphQL) *mux.Router {
	myRoute := mux.NewRouter()

	userController.InitialPath(myRoute, "/user")
	noteController.InitialPath(myRoute, "/note")

	myRoute.HandleFunc("/graphql", myGraphQLHandler.MyCustomGraphQLHandler(graphql.GetHandlerFunc().Schema))
	myRoute.Handle("/graphql-playground", myGraphQLHandler.Playground("gr", "/graphql")).Methods(http.MethodGet)

	myRoute.HandleFunc("/", checkAPI)
	return myRoute
}

func checkAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="localhost"`)

		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == "USERNAME") && (password == "PASSWORD")
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}
		//logic
		fmt.Println("ini authMiddleware")

		next.ServeHTTP(w, r)
	})
}
