package router

import (
	"fiber/fiber-graphql/config/database"
	"fiber/fiber-graphql/graphql/resolver"
	"fiber/fiber-graphql/graphql/schema"
	"fiber/fiber-graphql/repository"
	"fiber/fiber-graphql/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Gh *handler.Handler

func CustomerRouter(fiber *fiber.App, conf database.Config) {
	db := database.InitDB(conf)

	repo := repository.NewCustomerRepository(db)
	svc := service.NewCustomerService(repo, conf)

	cResolver := resolver.NewCustomerResolver(svc)
	cSchema := schema.NewCustomerSchema(cResolver)
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    cSchema.Query(),
		Mutation: cSchema.Mutation(),
	})

	if err != nil {
		log.Println(err)
	}

	Gh = handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})
}
