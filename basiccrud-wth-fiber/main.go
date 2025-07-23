package main

import (
	"cars/config"
	"cars/handlers"
	middleware "cars/middlewares"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(middleware.SecurityHeaders)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "12345",
			"manager": "qwerty",
			"john": "doe",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user is not authorized",
			})
		},
				
	}))

	app.Use(etag.New())

	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Put("/cars/:id", handlers.UpdateCar)

	
	fmt.Println("fiber http is listening on port 3015")
	// http.ListenAndServe(":3015", wrapperMux)
	if err := app.Listen(":3015"); err != nil {
		log.Fatal("couldnt listen on port 3015, error %v", err)
	}
}