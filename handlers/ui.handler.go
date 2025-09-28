package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func UIHandler(router fiber.Router) {
	var producer_ui = router.Group("/producers")
	var actor_ui = router.Group("/actors")
	var movie_ui = router.Group("/movies")

	producer_handler(producer_ui)
	actor_handler(actor_ui)
	movie_ui_handler(movie_ui)
}
