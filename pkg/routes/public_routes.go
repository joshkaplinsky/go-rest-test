package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshkaplinsky/go-rest-test/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	a.Delete("/", controllers.Delete)
	a.Get("/", controllers.Get)
	a.Patch("/", controllers.Patch)
	a.Post("/", controllers.Post)
	a.Put("/", controllers.Put)
}
