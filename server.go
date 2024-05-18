package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"vote.app/m/graph"
)

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx fiber.Ctx) {
	return func(ctx fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

func main() {
	app := fiber.New(fiber.Config{})
	// logger, _ := zap.NewProduction()

	// app.Use(service.New(service.Config{
	// 	Logger: logger,
	// }))
	// Create a gqlgen handler
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	h.AddTransport(transport.POST{})
	app.Post("/graphql", func(c fiber.Ctx) error {
		wrapHandler(h.ServeHTTP)(c)
		return nil
	})

	// Serve GraphQL Playground
	app.Get("/", func(c fiber.Ctx) error {
		wrapHandler(playground.Handler("GraphQL", "/graphql"))(c)
		return nil
	})

	// Start the server
	app.Listen(":3000", fiber.ListenConfig{
		DisableStartupMessage: false,
	})
}
