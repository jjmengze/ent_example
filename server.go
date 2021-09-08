package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github/jjmengze/test/graph"
	"github/jjmengze/test/graph/generated"
	"github/jjmengze/test/graph/model"
	"log"
	"net/http"
	"os"
	"strings"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	config := generated.Config{Resolvers: &graph.Resolver{}}
	config.Directives.Upper = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		switch obj.(type) {
		case *model.User:
			user := obj.(*model.User)
			UpperLastName := strings.ToUpper(*user.LastName)
			user.LastName = &UpperLastName

		}
		return next(ctx)
	}
	config.Directives.Date= func(ctx context.Context, obj interface{}, next graphql.Resolver, format *string) (res interface{}, err error) {
		fmt.Println(obj)
		fmt.Println(format)

		return next(ctx)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/",
		playground.Handler("GraphQL playground", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :" + port)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
