package handler

import (
	"context"
	"encoding/json"
	goGraphQL "github.com/graphql-go/graphql"
	goGraphQLHandler "github.com/graphql-go/handler"
	"net/http"
)

func New(schema *goGraphQL.Schema) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get query
		opts := goGraphQLHandler.NewRequestOptions(r)

		//rootValue := map[string]interface{}{
		//	"response": w,
		//	"request":  r,
		//	"viewer":   "john_doe",
		//}

		// execute graphql query
		params := goGraphQL.Params{
			Schema:         *schema,
			RequestString:  opts.Query,
			VariableValues: opts.Variables,
			OperationName:  opts.OperationName,
			Context:        context.WithValue(r.Context(), "Authorization", r.Header.Get("Authorization")),
		}

		//if graphHandler.rootObjectFn != nil {
		//	params.RootObject = h.rootObjectFn(ctx, r)
		//}

		result := goGraphQL.Do(params)

		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}
