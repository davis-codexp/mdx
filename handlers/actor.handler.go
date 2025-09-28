package handlers

import (
	"mdx/models"
	"mdx/utils"

	"github.com/gofiber/fiber/v2"
)

var actor_required_fields = []string{"Name", "Image"}

func actor_handler(router fiber.Router) {
	router.Post("/", add_actor)
}

func add_actor(ctx *fiber.Ctx) error {
	var branch = new(models.Actor)
	flag, _, message := utils.ParseAndValidate(branch, actor_required_fields, ctx)
	if flag {
		if branch.Add() {
			message = "Actor Added"
		} else {
			message = "Unable to Add Actor"
		}
	}
	return ctx.Status(fiber.StatusBadRequest).SendString(message)
}
