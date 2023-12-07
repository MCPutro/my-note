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

	userRouter(myRoute, userController)

	noteRouter(myRoute, noteController)

	graphqlRouter(myRoute, graphql)

	myRoute.HandleFunc("/", checkAPI)
	return myRoute
}

func userRouter(route *mux.Router, userController controller.UserController) {
	uRouter := route.PathPrefix("/user").Subrouter()
	uRouter.HandleFunc("/signUp", userController.CreateNewUser).Methods("POST")
	uRouter.HandleFunc("/signIn", userController.SignInUser).Methods("POST")
	uRouter.HandleFunc("/", userController.GetAllUser).Methods("GET")
}

func noteRouter(route *mux.Router, noteController controller.NoteController) {
	nRouter := route.PathPrefix("/note").Subrouter()
	nRouter.HandleFunc("/create", noteController.CreateNew).Methods("POST")
	nRouter.HandleFunc("/update", noteController.Update).Methods("POST")
	nRouter.HandleFunc("/getAllByUID", noteController.GetByUserId).Methods("GET")
	nRouter.HandleFunc("/remove", noteController.Delete).Methods("GET")
	nRouter.HandleFunc("/removePermanent", noteController.DeletePermanent).Methods("GET")
}

func graphqlRouter(route *mux.Router, graphql myGraphQL.GraphQL) {
	route.HandleFunc("/graphql", myGraphQLHandler.New(graphql.GetHandler().Schema))
	route.Handle("/graphql-playground", myGraphQLHandler.Playground("GraphQL Playground", "/graphql")).Methods(http.MethodGet)
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
