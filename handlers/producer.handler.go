package handlers

import (
	"mdx/models"
	"mdx/utils"

	"github.com/gofiber/fiber/v2"
)

var producer_required_fields = []string{"Name", "Mobile", "Address"}

func producer_handler(router fiber.Router) {
	router.Post("/", add_producer)
}

func add_producer(ctx *fiber.Ctx) error {
	var branch = new(models.Producer)
	flag, _, message := utils.ParseAndValidate(branch, producer_required_fields, ctx)
	if flag {
		if branch.Add() {
			message = "Producer Added"
		} else {
			message = "Unable to Add Producer"
		}
	}
	return ctx.Status(fiber.StatusBadRequest).SendString(message)
}
