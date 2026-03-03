package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	myTestApp := fiber.New()

	myTestApp.Get("/v1/tests", func(context *fiber.Ctx) error {
		testSample := Test2{"T-987894166519", "Sample of a good name", 36519819819514}

		return context.SendString(testSample.Id)
	})

	log.Println("HOLA")

	myTestApp.Listen(":8080")
}

type Test2 struct {
	Id          string `json: id`
	Name        string `json: name`
	createdTime int64  `json: time_created`
}
