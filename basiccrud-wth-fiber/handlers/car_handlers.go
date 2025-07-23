
package handlers

import (
	"cars/models"
	"fmt"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

func CreateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "incorrect input body",
			"details":  err.Error(),
		})
	}

	car.Insert()

	fmt.Println("Car saved to the inventory with thew id: ", car.ID)

	return c.Status(fiber.StatusCreated).JSON(car)
}

func UpdateCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car.ID = id

	if err := car.Get(); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car not found with the given id",
		})
	}

	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "incorrect input body",
			"details":  err.Error(),
		})
	}

	car.Update()

	fmt.Println("car updated with the id: ", id)

	return c.Status(fiber.StatusCreated).JSON(car)
}

func GetCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}
	
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car.ID = id

	if err := car.Get(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "car with the given id not found",
			"id": car.ID,
		})
	}

	fmt.Println("car found with the id:", id)

	return c.Status(fiber.StatusOK).JSON(car)
}

func DeleteCar(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	car := &models.Car{}

    id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	car.ID = id
    car.Delete()

	fmt.Println("car deleted with the id: ", id)

	return c.SendStatus(fiber.StatusNoContent)
}
