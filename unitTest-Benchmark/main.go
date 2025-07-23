package main

import (
	"cars/config"
	"cars/handlers"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "cars/docs"
)

// @title CarInventory
// @version 1.0
// @description This is an API to serve car inventory applications
// @host localhost:3015
func main() {

	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())
	// app.Use(middleware.SecurityHeaders)

	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "12345",
	// 		"manager": "qwerty",
	// 		"john": "doe",
	// 	},
	// 	Unauthorized: func(c *fiber.Ctx) error {
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 			"error": "user is not authorized",
	// 		})
	// 	},
				
	// }))

	app.Use(etag.New())
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Post("/cars", handlers.CreateCar)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Put("/cars/:id", handlers.UpdateCar)

	
	fmt.Println("fiber http is listening on port 3015")
	// http.ListenAndServe(":3015", wrapperMux)
	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("couldnt listen on port 3015, error %v", err)
	}
}