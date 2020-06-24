package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query string `json:"query"`
}

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

// New initial func started on main.go
func New() {
	charge()

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/api")
	if err != nil {
		panic(err)
	}

	if os.Getenv("PRODUCTION_MODE") == "false" {
		http.Handle("/graphiql", graphiqlHandler)
		fmt.Println("Acesse GRAPHIQL on http://caprisoft.com.br:4001/graphiql")
	}

	http.Handle("/api", gqlHandler())
	fmt.Println("Access the API on http://caprisoft.com.br:4001/api")
}

func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		fmt.Fprintf(w, "%s", processQuery(rBody.Query, schema))

	})
}

//Define the GraphQL Schema

func processQuery(query string, schema graphql.Schema) (result string) {

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		fmt.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)

	return fmt.Sprintf("%s", rJSON)
}
