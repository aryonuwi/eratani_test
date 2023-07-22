package routes

import (
	"github.com/aryonuwi/eratani_test/soaltiga/handlers"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(route *fiber.App) {
	route.Get("/country", handlers.GetCountry)
	route.Get("users/:country", handlers.GetUser)
	route.Post("users/", handlers.PostUser)
}
