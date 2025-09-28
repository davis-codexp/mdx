package handlers

import (
	"mdx/models"
	"mdx/utils"

	"github.com/gofiber/fiber/v2"
)

var movie_required_fields = []string{
	"Name", "Genre", "Length", "Year", "Synopsis", "Price",
	"Producer", "Thumbnail",
}

func movie_handler(router fiber.Router) {
	router.Get("/", get_movies)
	router.Get("/:id", get_movie_by_id)
}

func movie_ui_handler(router fiber.Router) {
	router.Post("/", add_movie)
}

/*----------------------API ROUTES-----------------*/
func get_movies(ctx *fiber.Ctx) error {
	var movie models.Movie
	var offset int = utils.GetIntQueryParams(ctx.Query("offset"), 0)
	var limit int = utils.GetIntQueryParams(ctx.Query("limit"), 10)
	var params = map[string]string{
		"name":  ctx.Query("name"),
		"genre": ctx.Query("genre"),
	}
	count, result, err := movie.Get(offset, limit, params)
	var count_data int = models.SanitizeInt(count.Count)
	return utils.GetArrayResponse(ctx, count_data, result, "Unable to get Movies", err)
}

func get_movie_by_id(ctx *fiber.Ctx) error {
	var movie models.Movie
	result, err := movie.GetById(ctx.Params("id"))
	return utils.GetSingleResponse(ctx, result, "Unable to get Movie", err)
}

/*----------------------UI ROUTES-----------------*/
func add_movie(ctx *fiber.Ctx) error {
	var movie = new(models.Movie)
	flag, _, message := utils.ParseAndValidate(movie, movie_required_fields, ctx)
	if flag {
		if movie.Add() {
			message = "Movie Added"
		} else {
			message = "Unable to Add Movie"
		}
	}
	return ctx.Status(fiber.StatusBadRequest).SendString(message)
}
