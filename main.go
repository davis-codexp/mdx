package main

import (
	"log"
	"mdx/handlers"
	"mdx/utils"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
	var PORT string = os.Getenv("PORT")
	conn, err := utils.GetConnection()
	if err != nil {
		log.Fatal("Unable to establish DB Connection", err.Error())
		panic(err.Error())
	}
	defer conn.Close()
	var app *fiber.App = fiber.New()
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	app.Use("/videos", filesystem.New(filesystem.Config{
		Root:   http.Dir("./videos"),
		Browse: false,
	}))
	var api_router fiber.Router = app.Group("/api/v1")
	var ui_router fiber.Router = app.Group("/")
	handlers.APIHandler(api_router)
	handlers.UIHandler(ui_router)
	if err = app.Listen(":" + PORT); err != nil {
		panic(err.Error())
	}
}
