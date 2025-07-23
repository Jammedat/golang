package handlers

import (
	"cars/config"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//github.com/stretchr/testify/assert

func TestCarAdd(t *testing.T) {
	config.ConnectDB()

	app := fiber.New()
	app.Post("/cars", CreateCar)

	body := `
	{
	"name": "corolla",
	"model": "xt",
	"brand": "toyota",
	"year": 2023,
	"price": 2873239
	
	}`

	req, _ := http.NewRequest("POST", "/cars", strings.NewReader(body))
	req.Header.Set("Content-Type", "application.json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)


}

func TestCarGet(t *testing.T) {
	config.ConnectDB()

	app := fiber.New()
	app.Get("/cars/:id", GetCar)


	req, _ := http.NewRequest("GET", "/cars/20", nil)
	req.Header.Set("Content-Type", "application.json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)


}
