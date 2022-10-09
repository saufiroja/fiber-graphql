package server

import (
	"fiber/fiber-graphql/config/database"
	"fiber/fiber-graphql/infrastructure/router"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func InitServer() *fiber.App {
	app := fiber.New()
	conf := database.Config{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GraphQL API Using Fiber")
	})

	router.CustomerRouter(app, conf)

	app.Get("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			router.Gh.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})

	app.Post("/graph", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			router.Gh.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})

	return app
}
