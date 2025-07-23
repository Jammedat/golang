package handlers

import (
	"cars/config"
	"net/http"

	"testing"

	"github.com/gofiber/fiber/v2"

)

//github.com/stretchr/testify/assert


func BenchmarkCarGet(b *testing.B) {
	config.ConnectDB()

	app := fiber.New()
	app.Get("/cars/:id", GetCar)


	req, _ := http.NewRequest("GET", "/cars/20", nil)
	req.Header.Set("Content-Type", "application.json")

	for i := 0; i < b.N; i++ {
		_, _ = app.Test(req, 5000)
	}


}
