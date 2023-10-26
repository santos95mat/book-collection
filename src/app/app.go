package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/santos95mat/go-book-collection/src/app/routes"
)

var app = fiber.New()
var v1 = app.Group("/v1")

func Run() {
	app.Use(cors.New())

	getRoutes()

	app.Listen(os.Getenv("PORT"))
}

func getRoutes() {
	routes.AddBookRoutes(v1)
	routes.AddStatusRoute(v1)
}
