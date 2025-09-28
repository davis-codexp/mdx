package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func APIHandler(router fiber.Router) {
	var movie_api = router.Group("/movies")

	movie_handler(movie_api)
}
