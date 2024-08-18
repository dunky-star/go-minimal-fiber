package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Get request
	app.Get("/", apiGet)

	// Get by parameter
	app.Get("/:param", apiGetParam)

	// Post data
	app.Post("/", apiPost)

	log.Fatal(app.Listen(":3000"))

}

func apiGet(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to KForce to in 2024!!!")
}

func apiGetParam(ctx *fiber.Ctx) error {
	return ctx.SendString("Param :" + ctx.Params("params"))
}

func apiPost(ctx *fiber.Ctx) error {
	return ctx.SendString("POST request")
}
